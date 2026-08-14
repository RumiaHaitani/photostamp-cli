package watermark

import (
    "image"
    "testing"
)

func BenchmarkApply(b *testing.B) {
    b.Log("Бенчмарк запущен!")
    src := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
    wm := image.NewRGBA(image.Rect(0, 0, 200, 100))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Apply(src, wm, 10, 1.0)
    }
}
