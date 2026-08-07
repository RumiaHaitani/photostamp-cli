package watermark

import (
	"image"
	"image/draw"
)

// Apply накладывает водяной знак wm на изображение src в правый нижний угол с отступом margin.
// Если scale != 1.0, знак масштабируется.
func Apply(src image.Image, wm image.Image, margin int, scale float64) image.Image {
	// Если нужно масштабирование, ресайзим водяной знак
	if scale != 1.0 {
		wm = resizeImage(wm, scale)
	}

	srcBounds := src.Bounds()
	wmBounds := wm.Bounds()
	x := srcBounds.Dx() - wmBounds.Dx() - margin
	y := srcBounds.Dy() - wmBounds.Dy() - margin
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	dst := image.NewRGBA(srcBounds)
	draw.Draw(dst, srcBounds, src, image.Point{0, 0}, draw.Src)
	draw.Draw(dst, wmBounds.Add(image.Point{x, y}), wm, image.Point{0, 0}, draw.Over)
	return dst
}

// resizeImage — простой ресайз с билинейной интерполяцией (грубо, но для demo сойдёт)
func resizeImage(img image.Image, scale float64) image.Image {
	if scale == 1.0 {
		return img
	}
	bounds := img.Bounds()
	newW := int(float64(bounds.Dx()) * scale)
	newH := int(float64(bounds.Dy()) * scale)
	if newW < 1 || newH < 1 {
		newW, newH = 1, 1
	}
	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))
	// Простая билинейная интерполяция
	for y := 0; y < newH; y++ {
		for x := 0; x < newW; x++ {
			srcX := float64(x) / float64(newW) * float64(bounds.Dx())
			srcY := float64(y) / float64(newH) * float64(bounds.Dy())
			// для простоты берём ближайший пиксель
			ix := int(srcX)
			iy := int(srcY)
			if ix >= bounds.Dx() {
				ix = bounds.Dx() - 1
			}
			if iy >= bounds.Dy() {
				iy = bounds.Dy() - 1
			}
			dst.Set(x, y, img.At(ix, iy))
		}
	}
	return dst
}
