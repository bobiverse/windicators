package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/bobiverse/windicators"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	// Define window size and position
	iw, err := windicators.NewIndicatorWindow(200, 20, windicators.PositionCenterTop)
	if err != nil {
		log.Fatalf("new window error: %s", err)
	}
	defer glfw.Terminate()

	// First indicator component
	windicators.NewComponent(iw, "demo-rand", "Demo(%v)", func(c *windicators.Component) any {
		r := rand.Intn(10)  // simulate some background job/api call
		c.IsVisible = r > 3 // when to show
		return r            // return value
	}, 2*time.Second)

	// Second indicator component
	windicators.NewComponent(iw, "demo-dummy-emails", "Unread->%v", func(c *windicators.Component) any {
		unreadCount := myFuncFetchEmails()
		c.IsVisible = unreadCount > 0
		return unreadCount
	}, 5*time.Second)

	// Ready!
	iw.Run()
}

// This is dummy function just to demonstrate something is done
func myFuncFetchEmails() int {
	return rand.Intn(5)
}
