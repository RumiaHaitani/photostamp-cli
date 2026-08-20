package camera

import (
    "bytes"
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
    log.Printf("Попытка открыть камеру с индексом %d", deviceID)
    webcam, err := gocv.OpenVideoCapture(deviceID)
    if err != nil {
        log.Printf("Ошибка открытия: %v", err)
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
    buf, err := gocv.IMEncode(".jpg", img)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования: %v", err)
    }
    defer buf.Close()
    reader := bytes.NewReader(buf.GetBytes())
    res, _, err := image.Decode(reader)
    if err != nil {
        return nil, fmt.Errorf("ошибка декодирования: %v", err)
    }
    return res, nil
}

func (d *GoCVDriver) Close() error {
    if d.webcam != nil {
        return d.webcam.Close()
    }
    return nil
}
