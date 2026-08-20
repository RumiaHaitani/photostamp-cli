package watermark

import (
    "image"
    "image/draw"

    xdraw "golang.org/x/image/draw"
)

func Apply(src image.Image, wm image.Image, margin int, scale float64) image.Image {
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
    xdraw.BiLinear.Scale(dst, dst.Bounds(), img, bounds, xdraw.Over, nil)
    return dst
}
