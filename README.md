```
# photostamp-cli

Консольная утилита для захвата фото с камеры и наложения полупрозрачного водяного знака (PNG).

## 🚀 Возможности

- Захват кадра с веб-камеры через GoCV (OpenCV) или из файла (для отладки).
- Наложение PNG-водяного знака с поддержкой альфа-канала и масштабированием.
- Работа полностью в памяти: без сохранения временных файлов на диск.
- CLI-интерфейс с гибкими параметрами.
- Бенчмарк для производительности наложения знака.

## 📦 Требования

- Go 1.21+
- OpenCV (устанавливается через GoCV, см. [инструкцию](https://gocv.io/getting-started/windows/))
- Для Windows: MinGW и CMake (для сборки OpenCV)

## 🛠️ Сборка

```bash
make build
```

Или на Windows (MinGW):

```bash
mingw32-make build
```

Бинарник появится в `bin/photostamp-cli`.

## 📸 Использование

```bash
./bin/photostamp-cli -driver gocv -watermark testdata/logo.png -margin 20 -scale 0.5
```

### Параметры

| Флаг | Описание | По умолчанию |
|------|----------|--------------|
| `-driver` | Драйвер: `dummy`, `file`, `gocv`, `win` | `dummy` |
| `-source` | Путь к файлу (для `file`) или ID камеры (для `gocv`) | `""` |
| `-watermark` | Путь к PNG-файлу водяного знака | `testdata/logo.png` |
| `-output` | Папка для сохранения результата | `output` |
| `-margin` | Отступ от правого нижнего края (пиксели) | `10` |
| `-scale` | Масштаб водяного знака (1.0 = оригинал) | `0.5` |

### Примеры

**Тестовый драйвер (генерирует градиент):**
```bash
./bin/photostamp-cli -driver dummy
```

**Чтение из файла:**
```bash
./bin/photostamp-cli -driver file -source testdata/sample.jpg
```

**Реальная камера:**
```bash
./bin/photostamp-cli -driver gocv -watermark mylogo.png -margin 30 -scale 0.7
```

## 🧪 Тесты и бенчмарки

```bash
make test
make bench
```

Или напрямую:
```bash
go test -v ./...
go test -bench=. -run=^$ ./internal/watermark
```

## 🏗️ Архитектура

Проект построен на интерфейсе `Camera` с различными реализациями (драйверами):
- `DummyCamera` – генерирует тестовое изображение.
- `FileCamera` – читает изображение с диска.
- `GoCVDriver` – захват с реальной камеры через OpenCV (Windows).

Все пакеты находятся в `internal/`, что предотвращает их импорт извне.

```