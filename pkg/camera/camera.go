package camera

import "image"

// Camera — интерфейс для драйвера камеры.
type Camera interface {
	Capture() (image.Image, error) // сделать снимок
	Close() error                  // освободить ресурсы
}

// New создаёт драйвер по имени.
func New(name string, source string) (Camera, error) {
	switch name {
	case "file":
		return NewFileCamera(source)
	case "dummy":
		return NewDummyCamera()
	default:
		return NewDummyCamera()
	}
}