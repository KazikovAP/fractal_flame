package fractal_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type mockFlameGenerator struct {
	*fractal.BaseFlameGenerator
}

func (fg *mockFlameGenerator) UpdatePixel(img *image.RGBA, point fractal.Point, c color.Color) {
	img.Set(int(point.X), int(point.Y), c)
}

func (fg *mockFlameGenerator) MapToCanvas(point fractal.Point, canvasWidth, canvasHeight int) fractal.Point {
	return fractal.Point{
		X: point.X * float64(canvasWidth) / 100.0,
		Y: point.Y * float64(canvasHeight) / 100.0,
	}
}

func (fg *mockFlameGenerator) RenderToImage(img *image.RGBA) error {
	bgColor := color.RGBA{255, 255, 255, 255}

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, bgColor)
		}
	}

	return nil
}

func TestUpdatePixel(t *testing.T) {
	flameGen := &mockFlameGenerator{
		BaseFlameGenerator: &fractal.BaseFlameGenerator{},
	}

	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	pixel := fractal.Point{X: 5.0, Y: 5.0}
	expectedColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	flameGen.UpdatePixel(img, pixel, expectedColor)

	actualColor := img.At(int(pixel.X), int(pixel.Y))
	if actualColor != expectedColor {
		t.Errorf("UpdatePixel failed, expected color %v, got %v", expectedColor, actualColor)
	}
}

func TestMapToCanvas(t *testing.T) {
	flameGen := &mockFlameGenerator{
		BaseFlameGenerator: &fractal.BaseFlameGenerator{},
	}

	point := fractal.Point{X: 50.0, Y: 75.0}
	canvasWidth, canvasHeight := 200, 200

	mappedPoint := flameGen.MapToCanvas(point, canvasWidth, canvasHeight)

	if mappedPoint.X < 0 || mappedPoint.X >= float64(canvasWidth) {
		t.Errorf("Mapped X coordinate out of bounds, expected 0 <= X < %d, got %f", canvasWidth, mappedPoint.X)
	}

	if mappedPoint.Y < 0 || mappedPoint.Y >= float64(canvasHeight) {
		t.Errorf("Mapped Y coordinate out of bounds, expected 0 <= Y < %d, got %f", canvasHeight, mappedPoint.Y)
	}
}

func TestRenderToImage(t *testing.T) {
	flameGen := &mockFlameGenerator{
		BaseFlameGenerator: &fractal.BaseFlameGenerator{},
	}

	width, height := 100, 100
	outputImage := image.NewRGBA(image.Rect(0, 0, width, height))

	err := flameGen.RenderToImage(outputImage)

	if err != nil {
		t.Errorf("RenderToImage failed with error: %v", err)
	}

	centerColor := outputImage.At(width/2, height/2)
	if centerColor == nil {
		t.Errorf("Pixel at center of the image is nil")
	}
}
