package application_test

import (
	"bytes"
	"testing"

	"github.com/KazikovAP/fractal_flame/config"
	"github.com/KazikovAP/fractal_flame/internal/application"
	"github.com/KazikovAP/fractal_flame/internal/domain/transformations"
)

type MockIOAdapter struct {
	OutputBuffer *bytes.Buffer
}

func (m *MockIOAdapter) Output(content string) {
	m.OutputBuffer.WriteString(content)
}

func TestApp_Start_Success(t *testing.T) {
	cfg := &config.Config{
		Width:               2000,
		Height:              2000,
		Iterations:          1000,
		TransformFn:         "spherical",
		TransformationCount: 3,
		Symmetry:            80,
		Gamma:               1.0,
		Mode:                "single",
		OutputDir:           t.TempDir(),
	}

	outputBuffer := &bytes.Buffer{}
	ioAdapter := &MockIOAdapter{OutputBuffer: outputBuffer}

	transformFn := transformations.TransformationType(cfg.TransformFn)

	validFunctions := map[transformations.TransformationType]bool{
		transformations.Bubble:     true,
		transformations.Sinusoidal: true,
		transformations.Spherical:  true,
		transformations.Polar:      true,
		transformations.Waves:      true,
	}

	if !validFunctions[transformFn] {
		t.Fatalf("Unknown transformation function: %s", cfg.TransformFn)
	}

	factory, err := transformations.NewTransformationFactory(transformFn, transformations.RandomColor)
	if err != nil {
		t.Fatalf("Expected no error creating TransformationFactory, but got: %v", err)
	}

	app := application.NewApp(cfg, ioAdapter, factory)

	err = app.Start()
	if err != nil {
		t.Fatalf("Expected no error starting app, but got: %v", err)
	}

	expectedOutput := "Запуск генерации изображения фрактального пламени с параметрами:"
	if !bytes.Contains(outputBuffer.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain %q, but got: %s", expectedOutput, outputBuffer.String())
	}
}
