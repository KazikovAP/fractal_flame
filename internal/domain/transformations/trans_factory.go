package transformations

import (
	"errors"
	"image/color"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type TransformationFactory struct {
	TransformFn string
	ColorFunc   func() color.Color
}

func NewTransformationFactory(transformFn string, colorFunc func() color.Color) (*TransformationFactory, error) {
	validFunctions := map[string]bool{
		"bubble":     true,
		"sinusoidal": true,
		"spherical":  true,
		"polar":      true,
		"waves":      true,
	}

	if !validFunctions[transformFn] {
		return nil, errors.New("unknown transformation function: " + transformFn)
	}

	return &TransformationFactory{
		TransformFn: transformFn,
		ColorFunc:   colorFunc,
	}, nil
}

func (tf *TransformationFactory) CreateTransformation() fractal.Transformation {
	transformationColor := tf.ColorFunc()

	switch tf.TransformFn {
	case "bubble":
		return NewBubbleTransformation(transformationColor)
	case "sinusoidal":
		return NewSinusoidalTransformation(transformationColor)
	case "spherical":
		return NewSphericalTransformation(transformationColor)
	case "polar":
		return NewPolarTransformation(transformationColor)
	case "waves":
		return NewWavesTransformation(transformationColor, 1.0, 1.0, 0.5, 0.5)
	}

	return nil
}
