package camera

import (
	"image"
	"log"
	"strconv"
)

type Camera interface {
	Capture() (image.Image, error)
	Close() error
}

func New(name string, source string) (Camera, error) {
	log.Printf("Создание драйвера: name=%s, source=%s", name, source)
	switch name {
	case "file":
		return NewFileCamera(source)
	case "dummy":
		return NewDummyCamera()
	case "gocv", "win":
		deviceID := 0
		if source != "" {
			if id, err := strconv.Atoi(source); err == nil {
				deviceID = id
			}
		}
		log.Printf("Попытка открыть камеру с индексом %d", deviceID)
		return NewGoCVDriver(deviceID)
	default:
		log.Printf("Неизвестный драйвер, используем dummy")
		return NewDummyCamera()
	}
}
