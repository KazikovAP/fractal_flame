package fractal

import (
	"crypto/rand"
	"math/big"
)

func RandomNumber(maxN int) int {
	maxBig := big.NewInt(int64(maxN))
	randNum, _ := rand.Int(rand.Reader, maxBig)

	return int(randNum.Int64())
}

func randomFloat64() float64 {
	return float64(RandomNumber(1<<53)) / float64(1<<53)
}

func Clamp(value int) uint8 {
	if value < 0 {
		return 0
	} else if value > 255 {
		return 255
	}

	return uint8(value)
}

func NewDefaultRect() Rect {
	return Rect{
		X:      -1,
		Y:      -1,
		Width:  2,
		Height: 2,
	}
}

func NewRandomPoint() Point {
	return Point{
		X: randomFloat64()*2 - 1,
		Y: randomFloat64()*2 - 1,
	}
}
