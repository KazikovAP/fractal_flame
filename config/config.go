package config

import (
	"errors"
	"flag"

	"golang.org/x/exp/constraints"
)

type Config struct {
	Width               int
	Height              int
	Iterations          int
	TransformFn         string
	TransformationCount int
	Symmetry            int
	Gamma               float64
	Mode                string
	OutputDir           string
}

type number interface {
	constraints.Integer | constraints.Float
}

func CorrectValue[T number](value, minValue, maxValue T) T {
	if value < minValue {
		return minValue
	} else if value > maxValue {
		return maxValue
	}

	return value
}

func Init() (*Config, error) {
	width := flag.Int("width", 2000, "Ширина изображения (по умолчанию 2000)")
	height := flag.Int("height", 2000, "Высота изображения (по умолчанию 2000)")
	iterations := flag.Int("iterations", 200000, "Количество итераций для генерации фрактала (по умолчанию 200000)")
	transformFn := flag.String("trans", "waves", "Функция трансформации (доступные: spherical, sinusoidal, bubble, polar, waves)")
	transformationCount := flag.Int("trans_count", 3, "Количество трансформаций (по умолчанию 3)")
	symmetry := flag.Int("symmetry", 80, "Число симметрий (по умолчанию 80)")
	gamma := flag.Float64("gamma", 1.0, "Гамма коррекция (по умолчанию 1.0)")
	mode := flag.String("mode", "single", "Режим выполнения программы (single - однопоточный или multi - многопоточный)")
	outputDir := flag.String("output_dir", "data", "Директория для сохранения результата (по умолчанию 'data')")

	flag.Parse()

	config := &Config{
		Width:               CorrectValue(*width, 3000, 5000),
		Height:              CorrectValue(*height, 3000, 5000),
		Iterations:          CorrectValue(*iterations, 100000, 250000),
		TransformFn:         *transformFn,
		TransformationCount: CorrectValue(*transformationCount, 1, 20),
		Symmetry:            CorrectValue(*symmetry, 20, 120),
		Gamma:               CorrectValue(*gamma, 1.0, 2.5),
		Mode:                *mode,
		OutputDir:           *outputDir,
	}

	if config.TransformFn == "" {
		return nil, errors.New("функция трансформации должна быть указана")
	}

	return config, nil
}
