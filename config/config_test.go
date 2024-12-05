package config_test

import (
	"flag"
	"testing"

	cfg "github.com/KazikovAP/fractal_flame/config"
	"github.com/stretchr/testify/assert"
)

func TestCorrectValue(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		minValue float64
		maxValue float64
		expected float64
	}{
		{"Value within range", 5.0, 1.0, 10.0, 5.0},
		{"Value below range", 0.5, 1.0, 10.0, 1.0},
		{"Value above range", 15.0, 1.0, 10.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cfg.CorrectValue(tt.value, tt.minValue, tt.maxValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInit(t *testing.T) {
	fs, config := cfg.DefineFlags()

	err := fs.Parse([]string{
		"-width=4000",
		"-height=5000",
		"-iterations=250000",
		"-trans=spherical",
		"-trans_count=10",
		"-symmetry=100",
		"-gamma=1.8",
		"-mode=multi",
	})
	assert.NoError(t, err)

	config.Width = cfg.CorrectValue(config.Width, 3000, 5000)
	config.Height = cfg.CorrectValue(config.Height, 3000, 5000)
	config.Iterations = cfg.CorrectValue(config.Iterations, 100000, 250000)
	config.TransformationCount = cfg.CorrectValue(config.TransformationCount, 1, 20)
	config.Symmetry = cfg.CorrectValue(config.Symmetry, 20, 120)
	config.Gamma = cfg.CorrectValue(config.Gamma, 1.0, 2.5)

	assert.Equal(t, 4000, config.Width)
	assert.Equal(t, 5000, config.Height)
	assert.Equal(t, 250000, config.Iterations)
	assert.Equal(t, "spherical", config.TransformFn)
	assert.Equal(t, 10, config.TransformationCount)
	assert.Equal(t, 100, config.Symmetry)
	assert.Equal(t, 1.8, config.Gamma)
	assert.Equal(t, "multi", config.Mode)
}

func TestInit_DefaultValues(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)

	_ = flag.CommandLine.Parse([]string{})

	config, err := cfg.Init()

	assert.NoError(t, err)
	assert.Equal(t, 3000, config.Width)
	assert.Equal(t, 3000, config.Height)
	assert.Equal(t, 200000, config.Iterations)
	assert.Equal(t, "waves", config.TransformFn)
	assert.Equal(t, 3, config.TransformationCount)
	assert.Equal(t, 80, config.Symmetry)
	assert.Equal(t, 1.0, config.Gamma)
	assert.Equal(t, "single", config.Mode)
}
