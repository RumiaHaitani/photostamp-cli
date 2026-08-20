package watermark

import (
    "image"
    "testing"
)

func BenchmarkApply(b *testing.B) {
    src := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
    wm := image.NewRGBA(image.Rect(0, 0, 200, 100))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Apply(src, wm, 10, 0.5)
    }
}

func TestApply(t *testing.T) {
    src := image.NewRGBA(image.Rect(0, 0, 100, 100))
    wm := image.NewRGBA(image.Rect(0, 0, 20, 20))
    result := Apply(src, wm, 10, 0.5)
    if result == nil {
        t.Error("Apply вернула nil")
    }
}
