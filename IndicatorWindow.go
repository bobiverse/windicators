package windicators

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/nullboundary/glfont"
	"golang.org/x/image/font/gofont/goregular"
)

// IndicatorWindow for indicators
type IndicatorWindow struct {
	*glfw.Window
	Font     *glfont.Font
	FontSize uint

	Position                  *Point
	Width, Height             uint
	ScreenWidth, ScreenHeight uint

	components ComponentList

	// internal channel to signal redraw only when necessary
	chRedraw chan bool
}

// NewIndicatorWindow ..
func NewIndicatorWindow(width, height, position uint) (*IndicatorWindow, error) {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.Floating, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.FocusOnShow, glfw.False)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	screenW, screenH := glfwScreenSize()

	w, err := glfw.CreateWindow(int(width), int(height), "Windicators", nil, nil)
	if err != nil {
		return nil, err
	}

	w.MakeContextCurrent()
	glfw.SwapInterval(1)
	w.SetOpacity(0)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	iw := &IndicatorWindow{
		Window:       w,
		Position:     PositionPoint(position, screenW, screenH, width, height),
		Width:        width,
		Height:       height,
		ScreenWidth:  screenW,
		ScreenHeight: screenH,
	}

	// load default font
	defaultFontPath := filepath.Join(os.TempDir(), "goregular.ttf") // make pull request to  `glfont` to include `LoadBytes`
	if err := os.WriteFile(defaultFontPath, goregular.TTF, 0600); err != nil {
		return nil, err
	}
	var fontSize uint = 12
	font, err := iw.UseFont(defaultFontPath, fontSize)
	if err != nil {
		return nil, err
	}
	font.SetColor(1, 1, 1, 1) // r,g,b,a font color

	// window position on screen
	iw.MoveTo(iw.Position.X, iw.Position.Y)

	iw.ListenEvents()
	return iw, nil
}

// Draw ..
func (iw *IndicatorWindow) Draw() {
	if !iw.IsRunning() {
		return
	}

	// No info, hide
	if iw.components.IsAllHidden() {
		// do not use `iw.Hide()/iw.Show()` as it focuses window and interrupts user
		iw.SetOpacity(0)
		iw.SetSize(1, 1)
		iw.SwapBuffers()
		return
	}

	// Clear for redraw
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Show default view with some active indicators
	if iw.GetOpacity() != 0.5 {
		iw.SetOpacity(0.5)
		iw.SetSize(int(iw.Width), int(iw.Height))
		iw.SwapBuffers()
	}

	// split available space in sectors and center each component inside
	oneComponentWidth := float32(iw.Width) / float32(len(iw.components))
	for ix, c := range iw.components {
		if !c.IsVisible {
			continue
		}
		x := float32(ix)*oneComponentWidth + (oneComponentWidth / 2) - float32(c.Width)/2
		y := float32(iw.Height / 2)
		c.RecalcFor(x, y, c.String())
		if c.IsValid() {
			c.DrawText(x, y)
		}
	}

	iw.SwapBuffers()
}

// IsRunning - if channel is active then window is started
func (iw *IndicatorWindow) IsRunning() bool {
	return iw != nil && iw.chRedraw != nil
}

// Run ..
func (iw *IndicatorWindow) Run() error {
	iw.chRedraw = make(chan bool)

	// No need for superfast recheck
	tickEvents := time.NewTicker(250 * time.Millisecond)

	// Trigger all components in background go-routine
	go func() {
		<-time.After(1 * time.Second) // wait so `for..select..` loop kicks in
		iw.components.CheckAll()
	}()

	for {
		select {

		case <-iw.chRedraw:
			iw.Draw()

		case <-tickEvents.C:
			glfw.PollEvents()
			if iw.ShouldClose() {
				return fmt.Errorf("should close")
			}
		}
	}
}

// MoveTo ..
func (iw *IndicatorWindow) MoveTo(x, y int) {
	iw.Window.SetPos(x, y)
}

// UseFont ..
func (iw *IndicatorWindow) UseFont(fpath string, size uint) (*glfont.Font, error) {
	font, err := glfont.LoadFont(fpath, int32(size), int(iw.Width), int(iw.Height))
	if err != nil {
		return nil, err
	}

	iw.Font = font
	iw.FontSize = size
	return iw.Font, nil
}

// ListenEvents ..
func (iw *IndicatorWindow) ListenEvents() {
	iw.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		// log.Printf("[IW] Event.. %v, %v ,%v", button, action, mods)
		isLeftClick := button == 0 && action == 1
		if !isLeftClick {
			return
		}

		posX, posY := iw.GetCursorPos()

		for _, c := range iw.components {
			// have callback defined
			if c.OnClick == nil {
				continue
			}

			// hitbox
			if !c.InPos(float32(posX), float32(posY)) {
				continue
			}

			// trigger component callback
			c.OnClick(c)
		}

	})
}

// Terminate window and clean up resources.
// Must be called whenever program exists - on success or on error.
// There is no need to call this function if IndicatorWindow creation fails.
func (iw *IndicatorWindow) Terminate() {
	// Must be called from main OS thread.
	runtime.LockOSThread()
	glfw.Terminate()
}
