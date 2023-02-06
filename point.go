package windicators

import (
	"fmt"
)

// Point ..
type Point struct {
	X, Y int
}

// NewPoint .. :shrug:
func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// String ..
func (p *Point) String() string {
	return fmt.Sprintf("%d;%d", p.X, p.Y)
}
