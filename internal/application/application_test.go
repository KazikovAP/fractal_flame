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
	}

	outputBuffer := &bytes.Buffer{}
	ioAdapter := &MockIOAdapter{OutputBuffer: outputBuffer}

	factory, _ := transformations.NewTransformationFactory(cfg.TransformFn, transformations.RandomColor)

	app := application.NewApp(cfg, ioAdapter, factory)

	err := app.Start()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	expectedOutput := "Запуск генерации изображения фрактального пламени с параметрами:"
	if !bytes.Contains(outputBuffer.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain %q, but got: %s", expectedOutput, outputBuffer.String())
	}
}
