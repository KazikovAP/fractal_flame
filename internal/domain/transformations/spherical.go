package transformations

import (
	"image/color"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type SphericalTransformation struct {
	Colorize
}

func (st *SphericalTransformation) Apply(p fractal.Point) fractal.Point {
	rSquared := p.X*p.X + p.Y*p.Y

	return fractal.Point{
		X: p.X / rSquared,
		Y: p.Y / rSquared,
	}
}

func NewSphericalTransformation(colorVal color.Color) *SphericalTransformation {
	return &SphericalTransformation{
		Colorize: Colorize{ColorVal: colorVal},
	}
}
