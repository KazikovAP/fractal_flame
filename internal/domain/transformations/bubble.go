package transformations

import (
	"image/color"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type BubbleTransformation struct {
	Colorize
}

func (bt *BubbleTransformation) Apply(p fractal.Point) fractal.Point {
	r2 := p.X*p.X + p.Y*p.Y
	factor := 4 / (4 + r2)

	return fractal.Point{
		X: p.X * factor,
		Y: p.Y * factor,
	}
}

func NewBubbleTransformation(colorVal color.Color) *BubbleTransformation {
	return &BubbleTransformation{
		Colorize: Colorize{ColorVal: colorVal},
	}
}
