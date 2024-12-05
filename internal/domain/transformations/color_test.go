package transformations_test

import (
	"image/color"
	"testing"

	ts "github.com/KazikovAP/fractal_flame/internal/domain/transformations"
)

func TestColorizeColor(t *testing.T) {
	c := &ts.Colorize{ColorVal: color.RGBA{R: 255, G: 0, B: 0, A: 255}}

	got := c.Color()
	want := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	if got != want {
		t.Errorf("Colorize.Color() = %v; want %v", got, want)
	}
}

func TestRandomColor(t *testing.T) {
	got := ts.RandomColor()

	if _, ok := got.(color.RGBA); !ok {
		t.Errorf("RandomColor() = %v; want a color.RGBA type", got)
	}

	r, g, b, a := got.RGBA()
	if r == 0 && g == 0 && b == 0 && a == 0 {
		t.Errorf("RandomColor() = %v; got a fully transparent color", got)
	}
}
