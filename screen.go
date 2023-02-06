package windicators

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// (c) https://stackoverflow.com/questions/65734201/how-to-get-width-and-height-of-monitor-in-python-glfw
func glfwScreenSize() (uint, uint) {
	monitor := glfw.GetPrimaryMonitor()
	vidmode := monitor.GetVideoMode()
	return uint(vidmode.Width), uint(vidmode.Height)
}
