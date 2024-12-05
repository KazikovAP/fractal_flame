package transformations

import (
	"image/color"
	"math"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type SinusoidalTransformation struct {
	Colorize
}

func (st *SinusoidalTransformation) Apply(p fractal.Point) fractal.Point {
	return fractal.Point{
		X: math.Sin(p.X),
		Y: math.Sin(p.Y),
	}
}

func NewSinusoidalTransformation(colorVal color.Color) *SinusoidalTransformation {
	return &SinusoidalTransformation{
		Colorize: Colorize{ColorVal: colorVal},
	}
}
