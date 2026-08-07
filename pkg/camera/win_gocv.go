package camera

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"

	"gocv.io/x/gocv"
)

type GoCVDriver struct {
	webcam *gocv.VideoCapture
}

func NewGoCVDriver(deviceID int) (*GoCVDriver, error) {
	log.Printf("Попытка открыть камеру с индексом %d", deviceID)
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть камеру %d: %v", deviceID, err)
	}
	webcam.Set(gocv.VideoCaptureFrameWidth, 640)
	webcam.Set(gocv.VideoCaptureFrameHeight, 480)
	time.Sleep(300 * time.Millisecond)
	log.Printf("Камера %d открыта", deviceID)
	return &GoCVDriver{webcam: webcam}, nil
}

func (d *GoCVDriver) Capture() (image.Image, error) {
	img := gocv.NewMat()
	defer img.Close()
	if !d.webcam.Read(&img) {
		return nil, fmt.Errorf("не удалось получить кадр с камеры")
	}
	if img.Empty() {
		return nil, fmt.Errorf("получен пустой кадр")
	}
	// Сохраняем кадр во временный файл
	tmpFile := "temp_capture.jpg"
	if ok := gocv.IMWrite(tmpFile, img); !ok {
		return nil, fmt.Errorf("не удалось сохранить временный файл")
	}
	// Читаем обратно как image.Image
	f, err := os.Open(tmpFile)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть временный файл: %v", err)
	}
	defer f.Close()
	res, err := jpeg.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования JPEG: %v", err)
	}
	// Удаляем временный файл (опционально)
	// os.Remove(tmpFile)
	return res, nil
}

func (d *GoCVDriver) Close() error {
	if d.webcam != nil {
		return d.webcam.Close()
	}
	return nil
}
