package application

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/KazikovAP/fractal_flame/config"
	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
	"github.com/KazikovAP/fractal_flame/internal/domain/transformations"
)

type ioAdapter interface {
	Output(content string)
}

type App struct {
	io      ioAdapter
	cfg     *config.Config
	factory *transformations.TransformationFactory
}

func NewApp(cfg *config.Config, io ioAdapter, factory *transformations.TransformationFactory) *App {
	return &App{cfg: cfg, io: io, factory: factory}
}

func (a *App) Start() error {
	a.io.Output("Запуск генерации изображения фрактального пламени с параметрами:\n")
	a.io.Output(fmt.Sprintf("Ширина: %d\n", a.cfg.Width))
	a.io.Output(fmt.Sprintf("Высота: %d\n", a.cfg.Height))
	a.io.Output(fmt.Sprintf("Итерации: %d\n", a.cfg.Iterations))
	a.io.Output(fmt.Sprintf("Функция трансформации: %s\n", a.cfg.TransformFn))
	a.io.Output(fmt.Sprintf("Количество трансформаций: %d\n", a.cfg.TransformationCount))
	a.io.Output(fmt.Sprintf("Число симметрий: %d\n", a.cfg.Symmetry))
	a.io.Output(fmt.Sprintf("Гамма-коррекция: %f\n", a.cfg.Gamma))
	a.io.Output(fmt.Sprintf("Режим выполнения: %s\n", a.cfg.Mode))

	transform := a.getTransformations(a.cfg.TransformationCount)

	var flameGen fractal.FlameGeneratorInterface

	if a.cfg.Mode == "multi" {
		flameGen = fractal.NewMultiFlameGenerator(a.cfg, transform)
	} else {
		flameGen = fractal.NewSingleFlameGenerator(a.cfg, transform)
	}

	var generatedImage *image.RGBA

	var elapsedTime time.Duration

	start := time.Now()

	generatedImage = flameGen.Generate(transform)

	elapsedTime = time.Since(start)
	a.io.Output(fmt.Sprintf("Время выполнения: %v\n", elapsedTime))

	outputDir := "data"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("ошибка создания директории %s: %w", outputDir, err)
	}

	outputFile := filepath.Join(outputDir, fmt.Sprintf("fractal_%s.png", a.cfg.TransformFn))

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, generatedImage); err != nil {
		return fmt.Errorf("ошибка сохранения изображения: %w", err)
	}

	a.io.Output(fmt.Sprintf("Фрактальное пламя сгенерировано и сохранено в файл: %s\n", outputFile))

	return nil
}

func (a *App) getTransformations(trCount int) []fractal.Transformation {
	transforms := make([]fractal.Transformation, trCount)
	for i := 0; i < trCount; i++ {
		transforms[i] = a.factory.CreateTransformation()
	}

	return transforms
}