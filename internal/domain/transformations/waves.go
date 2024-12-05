package transformations

import (
	"image/color"
	"math"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type WavesTransformation struct {
	Colorize
	FreqX float64
	FreqY float64
	AmpX  float64
	AmpY  float64
}

func (wt *WavesTransformation) Apply(p fractal.Point) fractal.Point {
	return fractal.Point{
		X: p.X + wt.AmpX*math.Sin(wt.FreqX*p.Y),
		Y: p.Y + wt.AmpY*math.Sin(wt.FreqY*p.X),
	}
}

func NewWavesTransformation(colorVal color.Color, freqX, freqY, ampX, ampY float64) *WavesTransformation {
	return &WavesTransformation{
		Colorize: Colorize{ColorVal: colorVal},
		FreqX:    freqX,
		FreqY:    freqY,
		AmpX:     ampX,
		AmpY:     ampY,
	}
}
