package fractal

import (
	"image"
	"image/color"
	"math"
)

type Canvas struct {
	Pixels [][]*Pixel
	Width  int
	Height int
}

func NewCanvas(width, height int) *Canvas {
	pixels := make([][]*Pixel, height)
	for i := range pixels {
		pixels[i] = make([]*Pixel, width)
		for j := range pixels[i] {
			pixels[i][j] = &Pixel{}
		}
	}

	return &Canvas{
		Pixels: pixels,
		Width:  width,
		Height: height,
	}
}

func (c *Canvas) UpdatePixel(x, y int, transformation Transformation) {
	if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
		pixel := c.Pixels[y][x]
		transformationColor := transformation.Color()

		if rgbaColor, ok := transformationColor.(color.RGBA); ok {
			pixel.R = Clamp(int(pixel.R)*pixel.HitCount + int(rgbaColor.R)/(pixel.HitCount+1))
			pixel.G = Clamp(int(pixel.G)*pixel.HitCount + int(rgbaColor.G)/(pixel.HitCount+1))
			pixel.B = Clamp(int(pixel.B)*pixel.HitCount + int(rgbaColor.B)/(pixel.HitCount+1))

			pixel.HitCount++
		}
	}
}

func (c *Canvas) MapToCanvas(world Rect, point Point) (x, y int) {
	x = int((point.X - world.X) * float64(c.Width) / world.Width)
	y = int((point.Y - world.Y) * float64(c.Height) / world.Height)

	return x, y
}

func (c *Canvas) RenderToImage(img *image.RGBA, gamma float64) {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			pixel := c.Pixels[y][x]
			img.Set(x, y, color.RGBA{R: pixel.R, G: pixel.G, B: pixel.B, A: 255})
		}
	}

	gammaCorrection(img, gamma)
}

func gammaCorrection(img *image.RGBA, gamma float64) {
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			img.Set(x, y, color.RGBA{
				R: uint8(math.Pow(float64(r)/255, gamma) * 255),
				G: uint8(math.Pow(float64(g)/255, gamma) * 255),
				B: uint8(math.Pow(float64(b)/255, gamma) * 255),
				A: Clamp(int(a >> 8)),
			})
		}
	}
}
