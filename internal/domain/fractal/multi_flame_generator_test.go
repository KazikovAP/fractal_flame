package fractal_test

import (
	"image/color"
	"runtime"
	"testing"

	"github.com/KazikovAP/fractal_flame/config"
	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
	"github.com/stretchr/testify/require"
)

type mockTransformation struct{}

func (m *mockTransformation) Apply(p fractal.Point) fractal.Point {
	return fractal.Point{X: p.X, Y: p.Y}
}

func (m *mockTransformation) Color() color.Color {
	return color.Black
}

func TestMultiFlameGenerator_Generate_LessThreadsThanCPU(t *testing.T) {
	workers := runtime.NumCPU() / 2

	cfg := &config.Config{
		Width:      3000,
		Height:     3000,
		Iterations: 1000,
	}

	transformations := []fractal.Transformation{&mockTransformation{}}
	generator := fractal.NewMultiFlameGenerator(cfg, transformations)

	generator.Workers = workers

	img := generator.Generate(transformations)

	require.Equal(t, cfg.Width, img.Bounds().Dx(), "Image width does not match the configuration")
	require.Equal(t, cfg.Height, img.Bounds().Dy(), "Image height does not match the configuration")
}

func TestMultiFlameGenerator_Generate_EqualThreadsAndCPU(t *testing.T) {
	workers := runtime.NumCPU()

	cfg := &config.Config{
		Width:      3000,
		Height:     3000,
		Iterations: 1000,
	}

	transformations := []fractal.Transformation{&mockTransformation{}}
	generator := fractal.NewMultiFlameGenerator(cfg, transformations)

	generator.Workers = workers

	img := generator.Generate(transformations)

	require.Equal(t, cfg.Width, img.Bounds().Dx(), "Image width does not match the configuration")
	require.Equal(t, cfg.Height, img.Bounds().Dy(), "Image height does not match the configuration")
}

func TestMultiFlameGenerator_Generate_MoreThreadsThanCPU(t *testing.T) {
	workers := runtime.NumCPU() * 2

	cfg := &config.Config{
		Width:      3000,
		Height:     3000,
		Iterations: 1000,
	}

	transformations := []fractal.Transformation{&mockTransformation{}}
	generator := fractal.NewMultiFlameGenerator(cfg, transformations)

	generator.Workers = workers

	img := generator.Generate(transformations)

	require.Equal(t, cfg.Width, img.Bounds().Dx(), "Image width does not match the configuration")
	require.Equal(t, cfg.Height, img.Bounds().Dy(), "Image height does not match the configuration")
}
