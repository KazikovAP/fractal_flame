package transformations

import (
	"image/color"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type Colorize struct {
	ColorVal color.Color
}

func (c *Colorize) Color() color.Color {
	return c.ColorVal
}

func RandomColor() color.Color {
	return color.RGBA{
		R: fractal.Clamp(fractal.RandomNumber(256)),
		G: fractal.Clamp(fractal.RandomNumber(256)),
		B: fractal.Clamp(fractal.RandomNumber(256)),
		A: 255,
	}
}
