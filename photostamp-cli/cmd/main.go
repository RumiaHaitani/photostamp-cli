package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/RumiaHaitani/photostamp-cli/pkg/camera"
	"github.com/RumiaHaitani/photostamp-cli/pkg/watermark"
)

func main() {
	// Параметры командной строки
	var (
		driver      = flag.String("driver", "dummy", "Драйвер: file, dummy")
		source      = flag.String("source", "", "Путь к файлу (для драйвера file)")
		watermarkFile = flag.String("watermark", "testdata/logo.png", "Путь к PNG с водяным знаком")
		outputDir   = flag.String("output", "output", "Папка для сохранения результата")
		margin      = flag.Int("margin", 10, "Отступ от края (пикселей)")
	)
	flag.Parse()

	// Создаём драйвер камеры
	cam, err := camera.New(*driver, *source)
	if err != nil {
		log.Fatalf("Ошибка инициализации драйвера: %v", err)
	}
	defer cam.Close()

	// Захват изображения
	img, err := cam.Capture()
	if err != nil {
		log.Fatalf("Ошибка захвата: %v", err)
	}

	// Наложение водяного знака
	result, err := watermark.Apply(img, *watermarkFile, *margin)
	if err != nil {
		log.Fatalf("Ошибка наложения водяного знака: %v", err)
	}

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