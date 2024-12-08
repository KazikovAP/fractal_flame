package transformations

import (
	"errors"
	"image/color"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
)

type TransformationType string

const (
	Bubble     TransformationType = "bubble"
	Sinusoidal TransformationType = "sinusoidal"
	Spherical  TransformationType = "spherical"
	Polar      TransformationType = "polar"
	Waves      TransformationType = "waves"
)

type TransformationFactory struct {
	TransformFn TransformationType
	ColorFunc   func() color.Color
}

func NewTransformationFactory(transformFn TransformationType, colorFunc func() color.Color) (*TransformationFactory, error) {
	validFunctions := map[TransformationType]bool{
		Bubble:     true,
		Sinusoidal: true,
		Spherical:  true,
		Polar:      true,
		Waves:      true,
	}

	if !validFunctions[transformFn] {
		return nil, errors.New("unknown transformation function: " + string(transformFn))
	}

	return &TransformationFactory{
		TransformFn: transformFn,
		ColorFunc:   colorFunc,
	}, nil
}

func (tf *TransformationFactory) CreateTransformation() fractal.Transformation {
	transformationColor := tf.ColorFunc()

	switch tf.TransformFn {
	case Bubble:
		return NewBubbleTransformation(transformationColor)
	case Sinusoidal:
		return NewSinusoidalTransformation(transformationColor)
	case Spherical:
		return NewSphericalTransformation(transformationColor)
	case Polar:
		return NewPolarTransformation(transformationColor)
	case Waves:
		return NewWavesTransformation(transformationColor, 1.0, 1.0, 0.5, 0.5)
	}

	return nil
}
