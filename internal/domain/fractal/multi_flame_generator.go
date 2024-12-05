package fractal

import (
	"image"
	"runtime"
	"sync"

	"github.com/KazikovAP/fractal_flame/config"
)

type MultiFlameGenerator struct {
	BaseFlameGenerator
}

func NewMultiFlameGenerator(cfg *config.Config, transformations []Transformation) *MultiFlameGenerator {
	return &MultiFlameGenerator{
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

func (fg *MultiFlameGenerator) Generate(transformations []Transformation) *image.RGBA {
	canvas := NewCanvas(fg.Width, fg.Height)
	world := NewDefaultRect()

	workers := runtime.NumCPU()
	wg := sync.WaitGroup{}
	jobChan := make(chan struct{}, workers)

	for w := 0; w < workers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for range jobChan {
				point := NewRandomPoint()
				fg.ProcessPoint(point, transformations, canvas, world)
			}
		}()
	}

	for i := 0; i < fg.Iterations; i++ {
		jobChan <- struct{}{}
	}

	close(jobChan)
	wg.Wait()

	img := image.NewRGBA(image.Rect(0, 0, fg.Width, fg.Height))
	canvas.RenderToImage(img, fg.Gamma)

	return img
}
