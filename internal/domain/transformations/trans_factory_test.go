package transformations_test

import (
	"image/color"
	"testing"

	ts "github.com/KazikovAP/fractal_flame/internal/domain/transformations"
)

func TestNewTransformationFactory(t *testing.T) {
	colorFunc := func() color.Color { return color.RGBA{R: 255, G: 0, B: 0, A: 255} }

	tf, err := ts.NewTransformationFactory("bubble", colorFunc)
	if err != nil {
		t.Errorf("NewTransformationFactory() error = %v; want no error", err)
	}

	if tf == nil {
		t.Error("NewTransformationFactory() = nil; want non-nil object")
	}

	_, err = ts.NewTransformationFactory("invalid", colorFunc)
	if err == nil {
		t.Error("NewTransformationFactory() error = nil; want error for unknown transformFn")
	}

	expectedErr := "unknown transformation function: invalid"
	if err.Error() != expectedErr {
		t.Errorf("NewTransformationFactory() error = %v; want %v", err.Error(), expectedErr)
	}
}

func TestCreateTransformation(t *testing.T) {
	colorFunc := func() color.Color { return color.RGBA{R: 255, G: 0, B: 0, A: 255} }

	tests := []struct {
		transformFn string
		expectError bool
	}{
		{"bubble", false},
		{"sinusoidal", false},
		{"spherical", false},
		{"polar", false},
		{"waves", false},
		{"invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.transformFn, func(t *testing.T) {
			tf, err := ts.NewTransformationFactory(tt.transformFn, colorFunc)
			if tt.expectError && err == nil {
				t.Errorf("expected error for transformFn %s, but got nil", tt.transformFn)
			}

			if !tt.expectError && err != nil {
				t.Errorf("did not expect error for transformFn %s, but got %v", tt.transformFn, err)
			}

			if !tt.expectError && tf != nil {
				transformation := tf.CreateTransformation()
				if transformation == nil {
					t.Errorf("CreateTransformation() returned nil for transformFn %s", tt.transformFn)
				}
			}
		})
	}
}
