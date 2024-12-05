package config_test

import (
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
