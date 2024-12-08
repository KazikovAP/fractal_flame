package transformations_test

import (
	"image/color"
	"testing"

	"github.com/KazikovAP/fractal_flame/internal/domain/fractal"
	ts "github.com/KazikovAP/fractal_flame/internal/domain/transformations"
	"github.com/stretchr/testify/require"
)

func TestNewTransformationFactory_ValidFunctions(t *testing.T) {
	colorFunc := func() color.Color { return color.RGBA{R: 255, G: 0, B: 0, A: 255} }

	tests := []struct {
		transformFn ts.TransformationType
		expected    fractal.Transformation
	}{
		{ts.Bubble, ts.NewBubbleTransformation(colorFunc())},
		{ts.Sinusoidal, ts.NewSinusoidalTransformation(colorFunc())},
		{ts.Spherical, ts.NewSphericalTransformation(colorFunc())},
		{ts.Polar, ts.NewPolarTransformation(colorFunc())},
		{ts.Waves, ts.NewWavesTransformation(colorFunc(), 1.0, 1.0, 0.5, 0.5)},
	}

	for _, tt := range tests {
		t.Run(string(tt.transformFn), func(t *testing.T) {
			factory, err := ts.NewTransformationFactory(tt.transformFn, colorFunc)
			require.NoError(t, err, "NewTransformationFactory() should not return an error for valid transformFn")
			require.NotNil(t, factory, "NewTransformationFactory() should return a non-nil factory")

			transformation := factory.CreateTransformation()
			require.EqualValues(t, tt.expected, transformation, "Created transformation does not match expected")
		})
	}
}

func TestNewTransformationFactory_InvalidFunction(t *testing.T) {
	colorFunc := func() color.Color { return color.RGBA{R: 255, G: 0, B: 0, A: 255} }

	_, err := ts.NewTransformationFactory("invalid", colorFunc)

	require.EqualError(t, err, "unknown transformation function: invalid",
		"Error message does not match expected")
}
