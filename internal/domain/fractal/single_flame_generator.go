package fractal

import (
	"image"

	"github.com/KazikovAP/fractal_flame/config"
)

type SingleFlameGenerator struct {
	BaseFlameGenerator
}

func NewSingleFlameGenerator(cfg *config.Config, transformations []Transformation) *SingleFlameGenerator {
	return &SingleFlameGenerator{
		BaseFlameGenerator: BaseFlameGenerator{
			Width:           cfg.Width,
			Height:          cfg.Height,
			Iterations:      cfg.Iterations,
			Transformations: transformations,
			Symmetry:        cfg.Symmetry,
			Gamma:           cfg.Gamma,
		},
	}
}

func (fg *SingleFlameGenerator) Generate(transformations []Transformation) *image.RGBA {
	canvas := NewCanvas(fg.Width, fg.Height)
	world := NewDefaultRect()

	for i := 0; i < fg.Iterations; i++ {
		point := NewRandomPoint()

		fg.ProcessPoint(point, transformations, canvas, world)
	}

	img := image.NewRGBA(image.Rect(0, 0, fg.Width, fg.Height))
	canvas.RenderToImage(img, fg.Gamma)

	return img
}
