package fractal

import (
	"image"
	"math"
)

type FlameGeneratorInterface interface {
	Generate(transformations []Transformation) *image.RGBA
}

type BaseFlameGenerator struct {
	Width           int
	Height          int
	Iterations      int
	Transformations []Transformation
	Symmetry        int
	Gamma           float64
}

func (fg *BaseFlameGenerator) ProcessPoint(point Point, transformations []Transformation, canvas *Canvas, world Rect) {
	transformation := transformations[RandomNumber(len(transformations))]
	point = transformation.Apply(point)

	if !world.contains(point) {
		return
	}

	for s := 0; s < fg.Symmetry; s++ {
		theta := float64(s) * 2 * math.Pi / float64(fg.Symmetry)
		rotatedPoint := point.rotate(theta)

		x, y := canvas.MapToCanvas(world, rotatedPoint)
		canvas.UpdatePixel(x, y, transformation)
	}
}
