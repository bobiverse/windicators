package windicators

import (
	"fmt"
	"log"
)

// ComponentMetrics ,,
type ComponentMetrics struct {
	parentComponent *Component

	X, Y   float32
	Height uint

	// calculated
	FontSize uint
	Baseline uint
	Width    uint
	X2, Y2   float32
}

// NewMetrics ..
func NewMetrics(c *Component) *ComponentMetrics {
	metr := &ComponentMetrics{
		parentComponent: c,
	}
	return metr
}

// IsValid ..
func (metr *ComponentMetrics) IsValid() bool {
	return metr != nil && metr.parentComponent != nil && metr.FontSize > 0 && metr.Baseline > 0 && metr.X2 >= 0
}

// RecalcFor ..
func (metr *ComponentMetrics) RecalcFor(posX, posY float32, s string) {
	metr.Width = uint(metr.parentComponent.Font.Width(1, s))
	metr.FontSize = metr.Height
	metr.Baseline = metr.FontSize - (metr.FontSize / 3)
	metr.X = posX
	metr.Y = posY - float32(metr.Baseline)
	metr.X2 = metr.X + float32(metr.Width)
	metr.Y2 = metr.Y + float32(metr.Height)

	if !metr.IsValid() {
		log.Printf("WARNING: Invalid metrics for %s", metr.parentComponent.ID)
	}
	// log.Println(metr.String())
}

// CoordinatesWithPadding - top, right, bottom, left padding
func (metr *ComponentMetrics) CoordinatesWithPadding(t, r, b, l float32) (float32, float32, float32, float32) {
	x := metr.X - l
	y := metr.Y - t
	x2 := metr.X2 + r
	y2 := metr.Y2 + b

	return x, y, x2, y2
}

// String ..
func (metr *ComponentMetrics) String() string {
	s := fmt.Sprintf("F[%d,%d]", metr.FontSize, metr.Baseline)
	s += fmt.Sprintf("\tXY[%.0f:%.0f]", metr.X, metr.Y)
	s += fmt.Sprintf("\tXY2[%.0f:%.0f]", metr.X2, metr.Y2)
	s += fmt.Sprintf("\tWH[%d;%d]", metr.Width, metr.Height)
	return s
}
