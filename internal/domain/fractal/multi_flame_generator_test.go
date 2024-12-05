package fractal_test

import (
	"image/color"
	"testing"

	"github.com/KazikovAP/fractal_flame/config"
	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type mockTransformation struct{}

func (m *mockTransformation) Apply(p fractal.Point) fractal.Point {
	return fractal.Point{X: p.X, Y: p.Y}
}

func (m *mockTransformation) Color() color.Color {
	return color.Black
}

func TestMultiFlameGenerator_Generate(t *testing.T) {
	cfg := &config.Config{
		Width:      3000,
		Height:     3000,
		Iterations: 1000,
	}

	transformations := []fractal.Transformation{&mockTransformation{}}
	generator := fractal.NewMultiFlameGenerator(cfg, transformations)

	img := generator.Generate(transformations)

	if img.Bounds().Dx() != cfg.Width || img.Bounds().Dy() != cfg.Height {
		t.Errorf("Generated image has wrong dimensions, expected (%d, %d), got (%d, %d)",
			cfg.Width, cfg.Height, img.Bounds().Dx(), img.Bounds().Dy())
	}
}
