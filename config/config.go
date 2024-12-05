package config

import (
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

func DefineFlags() (*flag.FlagSet, *Config) {
	fs := flag.NewFlagSet("config", flag.ExitOnError)
	config := &Config{}

	fs.IntVar(&config.Width, "width", 3000, "Ширина изображения (по умолчанию 3000)")
	fs.IntVar(&config.Height, "height", 3000, "Высота изображения (по умолчанию 3000)")
	fs.IntVar(&config.Iterations, "iterations", 200000, "Количество итераций для генерации фрактала (по умолчанию 200000)")
	fs.StringVar(&config.TransformFn, "trans", "waves", "Функция трансформации (доступные: spherical, sinusoidal, bubble, polar, waves)")
	fs.IntVar(&config.TransformationCount, "trans_count", 3, "Количество трансформаций (по умолчанию 3)")
	fs.IntVar(&config.Symmetry, "symmetry", 80, "Число симметрий (по умолчанию 80)")
	fs.Float64Var(&config.Gamma, "gamma", 1.0, "Гамма коррекция (по умолчанию 1.0)")
	fs.StringVar(&config.Mode, "mode", "single", "Режим выполнения программы (single - однопоточный или multi - многопоточный)")

	return fs, config
}

func Init() (*Config, error) {
	fs, config := DefineFlags()
	if err := fs.Parse(flag.Args()); err != nil {
		return nil, err
	}

	config.Width = CorrectValue(config.Width, 3000, 5000)
	config.Height = CorrectValue(config.Height, 3000, 5000)
	config.Iterations = CorrectValue(config.Iterations, 100000, 250000)
	config.TransformationCount = CorrectValue(config.TransformationCount, 1, 20)
	config.Symmetry = CorrectValue(config.Symmetry, 20, 120)
	config.Gamma = CorrectValue(config.Gamma, 1.0, 2.5)

	return config, nil
}
