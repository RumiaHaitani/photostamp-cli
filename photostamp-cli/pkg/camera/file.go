package camera

import (
	"image"
	_ "image/jpeg" // для поддержки JPEG
	_ "image/png"  // для поддержки PNG
	"os"
)

type FileCamera struct {
	path string
}

func NewFileCamera(path string) (*FileCamera, error) {
	return &FileCamera{path: path}, nil
}

func (c *FileCamera) Capture() (image.Image, error) {
	f, err := os.Open(c.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}

func (c *FileCamera) Close() error {
	return nil
}