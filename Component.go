package windicators

import (
	"fmt"
	"log"
	"time"

	"github.com/nullboundary/glfont"
)

// Component ..
type Component struct {
	parentWindow *IndicatorWindow

	ID          string
	Format      string
	Value       any
	IsVisible   bool
	fnCheck     func(c *Component) any
	fnCheckTick *time.Ticker

	URL string

	Font glfont.Font
	*ComponentMetrics

	OnClick func(c *Component)
}

// NewComponent ..
func NewComponent(iw *IndicatorWindow, id, format string, fnCheck func(c *Component) any, fnCheckDuration time.Duration) *Component {
	c := &Component{
		parentWindow: iw,
		ID:           id,
		Format:       format,
		fnCheck:      fnCheck,
		fnCheckTick:  time.NewTicker(fnCheckDuration),
	}

	if iw == nil || id == "" || format == "" || fnCheck == nil {
		return nil
	}

	c.Font = *iw.Font // copy, to allow every component different font settings

	c.ComponentMetrics = NewMetrics(c)
	c.Height = iw.FontSize

	// regular run
	go func(c *Component) {
		for {
			<-c.fnCheckTick.C
			if c.parentWindow.IsRunning() {
				c.Check()
			}
		}
	}(c)

	iw.components = append(iw.components, c)
	return c
}

// IsClickable have coordinates updated after draw
func (c *Component) IsClickable() bool {
	return c.ComponentMetrics.IsValid()
}

// Check ..
func (c *Component) Check() {
	oldValue := c.Value
	c.Value = c.fnCheck(c)
	// log.Printf("[%s] Check.. [%v] ==> [%v]", c.ID, oldValue, c.Value)

	// must redraw if changed
	if c.Value != oldValue {
		c.parentWindow.chRedraw <- true
	}
}

// DrawText ..
func (c *Component) DrawText(posX, posY float32) {
	posY += float32(c.FontSize - c.Baseline)
	_ = c.Font.Printf(posX, posY, 1.0, c.String()) // x,y,scale,string,printf args

	// posY IS NOT c.Y as font is drawn differently
	c.ComponentMetrics.RecalcFor(posX, posY, c.String())
}

// UseFont - change font for component
func (c *Component) UseFont(fpath string, pt int32) {
	font, err := glfont.LoadFont(fpath, pt, int(c.parentWindow.Width), int(c.parentWindow.Height))
	if err != nil {
		log.Println(err)
	}
	c.Height = uint(pt)
	c.FontSize = c.Height
	c.Font = *font
}

// String ..
func (c *Component) String() string {
	return fmt.Sprintf(c.Format, c.Value)
}

// InPos ..
func (c *Component) InPos(posX, posY float32) bool {
	if !c.IsVisible || !c.IsValid() {
		return false
	}

	x, y, x2, y2 := c.CoordinatesWithPadding(2, 10, 2, 10)

	hitX := posX >= x && posX <= x2
	hitY := posY >= y && posY <= y2
	return hitX && hitY
}
