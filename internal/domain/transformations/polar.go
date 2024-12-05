package transformations

import (
	"image/color"
	"math"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type PolarTransformation struct {
	Colorize
}

func (pt *PolarTransformation) Apply(p fractal.Point) fractal.Point {
	radius := math.Sqrt(p.X*p.X + p.Y*p.Y)
	theta := math.Atan2(p.Y, p.X)

	return fractal.Point{
		X: theta / math.Pi,
		Y: radius - 1.0,
	}
}

func NewPolarTransformation(colorVal color.Color) *PolarTransformation {
	return &PolarTransformation{
		Colorize: Colorize{ColorVal: colorVal},
	}
}
