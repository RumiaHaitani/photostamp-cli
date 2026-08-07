package camera

import (
	"image"
	"image/color"
)

type DummyCamera struct{}

func NewDummyCamera() (*DummyCamera, error) {
	return &DummyCamera{}, nil
}

func (c *DummyCamera) Capture() (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))
	for y := 0; y < 480; y++ {
		for x := 0; x < 640; x++ {
			r := uint8(x * 255 / 640)
			g := uint8(y * 255 / 480)
			b := uint8((x + y) * 255 / (640 + 480))
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	return img, nil
}

func (c *DummyCamera) Close() error {
	return nil
}
