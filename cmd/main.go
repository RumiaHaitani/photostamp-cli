package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"photostamp-cli/internal/camera"
	"photostamp-cli/internal/watermark"
	"time"
)

// loadWatermark загружает PNG из файла один раз при старте
func loadWatermark(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func main() {
	// Параметры командной строки
	var (
		driver        = flag.String("driver", "dummy", "Драйвер: file, dummy, gocv, win")
		source        = flag.String("source", "", "Путь к файлу (для драйвера file) или ID камеры (для gocv)")
		watermarkFile = flag.String("watermark", "testdata/logo.png", "Путь к PNG с водяным знаком")
		outputDir     = flag.String("output", "output", "Папка для сохранения результата")
		margin        = flag.Int("margin", 10, "Отступ от края (пикселей)")
		scale         = flag.Float64("scale", 0.5, "Масштаб водяного знака (1.0 = оригинал, 0.5 = вдвое меньше)")
	)
	flag.Parse()

	log.Printf("Выбран драйвер: %s", *driver)

	// Загружаем водяной знак один раз при старте
	watermarkImg, err := loadWatermark(*watermarkFile)
	if err != nil {
		log.Fatalf("Не удалось загрузить водяной знак: %v", err)
	}

	// Создаём драйвер камеры
	cam, err := camera.New(*driver, *source)
	if err != nil {
		log.Fatalf("Ошибка инициализации драйвера: %v", err)
	}
	log.Printf("Драйвер создан успешно")
	defer cam.Close()

	// Захват изображения
	img, err := cam.Capture()
	if err != nil {
		log.Fatalf("Ошибка захвата: %v", err)
	}

	// Наложение водяного знака (Apply теперь не возвращает ошибку)
	result := watermark.Apply(img, watermarkImg, *margin, *scale)

	// Создаём папку output, если её нет
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Не могу создать папку %s: %v", *outputDir, err)
	}

	// Имя файла с timestamp
	filename := fmt.Sprintf("photo_%s.jpg", time.Now().Format("20060102_150405"))
	outPath := filepath.Join(*outputDir, filename)

	// Сохраняем как JPEG
	outFile, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("Не могу создать файл: %v", err)
	}
	defer outFile.Close()

	if err := jpeg.Encode(outFile, result, &jpeg.Options{Quality: 95}); err != nil {
		log.Fatalf("Ошибка сохранения JPEG: %v", err)
	}

	fmt.Printf("✅ Фото сохранено: %s\n", outPath)
}
