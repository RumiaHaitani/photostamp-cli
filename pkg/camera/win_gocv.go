package camera

import (
	"fmt"
	"image"
	"log"
	"time"

	"gocv.io/x/gocv"
)

type GoCVDriver struct {
	webcam *gocv.VideoCapture
}

func NewGoCVDriver(deviceID int) (*GoCVDriver, error) {
	log.Printf("▶ Открываем камеру с индексом %d", deviceID)
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		log.Printf("❌ Ошибка OpenVideoCapture: %v", err)
		return nil, fmt.Errorf("не удалось открыть камеру %d: %v", deviceID, err)
	}
	// Проверяем, что камера реально открыта, пытаясь прочитать один раз
	// Это часто выявляет проблемы
	testMat := gocv.NewMat()
	defer testMat.Close()
	if !webcam.Read(&testMat) {
		log.Printf("❌ Не удалось прочитать тестовый кадр при открытии")
		webcam.Close()
		return nil, fmt.Errorf("камера %d не даёт кадр", deviceID)
	}
	if testMat.Empty() {
		log.Printf("❌ Тестовый кадр пуст")
		webcam.Close()
		return nil, fmt.Errorf("камера %d вернула пустой кадр", deviceID)
	}
	log.Printf("✅ Камера %d успешно открыта, тестовый кадр получен (%dx%d)", deviceID, testMat.Cols(), testMat.Rows())
	// Устанавливаем разрешение (если поддерживается)
	webcam.Set(gocv.VideoCaptureFrameWidth, 640)
	webcam.Set(gocv.VideoCaptureFrameHeight, 480)
	time.Sleep(200 * time.Millisecond) // даём камере стабилизироваться
	return &GoCVDriver{webcam: webcam}, nil
}

func (d *GoCVDriver) Capture() (image.Image, error) {
	img := gocv.NewMat()
	defer img.Close()
	if !d.webcam.Read(&img) {
		log.Printf("❌ Capture: Read вернул false")
		return nil, fmt.Errorf("не удалось получить кадр")
	}
	log.Printf("📸 Capture: получен кадр %dx%d, каналов %d", img.Cols(), img.Rows(), img.Channels())
	if img.Empty() {
		log.Printf("❌ Capture: кадр пуст")
		return nil, fmt.Errorf("получен пустой кадр")
	}
	res, err := img.ToImage()
	if err != nil {
		log.Printf("❌ Capture: ToImage ошибка: %v", err)
		return nil, fmt.Errorf("ошибка конвертации: %v", err)
	}
	log.Printf("✅ Capture: успешно конвертирован")
	return res, nil
}

func (d *GoCVDriver) Close() error {
	if d.webcam != nil {
		return d.webcam.Close()
	}
	return nil
}
