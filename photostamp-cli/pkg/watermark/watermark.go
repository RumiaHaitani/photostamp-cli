package watermark

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

// Apply накладывает полупрозрачный PNG (watermarkPath) на изображение src
// в правый нижний угол с отступом margin (пикселей).
func Apply(src image.Image, watermarkPath string, margin int) (image.Image, error) {
	// Загружаем водяной знак
	f, err := os.Open(watermarkPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	wm, err := png.Decode(f)
	if err != nil {
		return nil, err
	}

	// Размеры
	srcBounds := src.Bounds()
	wmBounds := wm.Bounds()

	// Вычисляем позицию (правый нижний угол)
	x := srcBounds.Dx() - wmBounds.Dx() - margin
	y := srcBounds.Dy() - wmBounds.Dy() - margin
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	// Создаём новый RGBA, копируем фон
	dst := image.NewRGBA(srcBounds)
	draw.Draw(dst, srcBounds, src, image.Point{0, 0}, draw.Src)

	// Накладываем водяной знак с прозрачностью (draw.Over)
	draw.Draw(dst, wmBounds.Add(image.Point{x, y}), wm, image.Point{0, 0}, draw.Over)

	return dst, nil
}