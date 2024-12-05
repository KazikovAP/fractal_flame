package fractal

import (
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type Rect struct {
	X, Y   float64
	Width  float64
	Height float64
}

type Pixel struct {
	R, G, B  uint8
	HitCount int
}

type Transformation interface {
	Apply(p Point) Point
	Color() color.Color
}

func (p Point) rotate(theta float64) Point {
	return Point{
		X: p.X*math.Cos(theta) - p.Y*math.Sin(theta),
		Y: p.X*math.Sin(theta) + p.Y*math.Cos(theta),
	}
}

func (p Point) inRect(r Rect) bool {
	return p.X >= r.X && p.X <= r.X+r.Width && p.Y >= r.Y && p.Y <= r.Y+r.Height
}

func (r Rect) contains(p Point) bool {
	return p.inRect(r)
}
