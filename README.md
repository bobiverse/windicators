# windicators
Simple GUI indicators for desktop (discrete, no blinking).  
Hidden until one of the indicators are valid for display.  
Use your own _jobs_: API calls, bash scripts etc. for component data source.

_Tested on Ubuntu_.

#### Example
```go
func main() {
	// Define window size and position
    iw, _ := wi.NewIndicatorWindow(200, 20, wi.PositionCenterBottom)
    defer glfw.Terminate()
    
	// First indicator component 
    wi.NewComponent(iw, "demo-rand", "Demo(%v)", func (c *wi.Component) any {
        r := rand.Intn(10) // simulate some background job/api call
        c.IsVisible = r > 3 // when to show
        return r            // return value
    }, 2*time.Second)
    	
	// Second indicator component 
    wi.NewComponent(iw, "demo-dummy-emails", "Unread->%v", func (c *wi.Component) any {
        unreadCount := myFuncFetchEmails()
        c.IsVisible = unreadCount > 0
        return unreadCount
    }, 5*time.Second)
    
    // Ready!
    wi.Run()
}

// This is dummy function just to demonstrate something is done 
func myFuncFetchEmails() int {
	 return rand.Intn(5)
}
```
![](demo.png)

#### Text colors
```go
	c := wi.NewComponent(iw, "demo-dummy-emails", "Unread->%v", func(c *wi.Component) any {
		unreadCount := myFuncFetchEmails()
		c.IsVisible = unreadCount > 0
		return unreadCount
	}, 5*time.Second)

	c.Font.SetColor(.2, 1.0, .4, 1)
	c.OnClick = func(c *wi.Component) {
		log.Printf("[%s] Click callback...", c.ID)
	}

```
![](demo2.png)

## Use `glfw.Window` and customize as you will
```go
// IndicatorWindow for indicators
type IndicatorWindow struct {
	*glfw.Window
	// ...
```

## Example to dynamically change colors
```go
	wi.NewComponent(iw, "demo-rand-status", "Status=%v", func(c *wi.Component) any {
		// simulate some background job/api call
		r := rand.Intn(5)

		status := ""
		switch r {
		case 1:
			status = "OK"
			c.Font.SetColor(0, 1, 0, 1)
		case 3:
			status = "ERR"
			c.Font.SetColor(1, 0, 0, 1)
		default:
			c.Font.SetColor(1, 1, 1, .5)
		}

		c.IsVisible = status != "" // when to show
		return status              // return value
	}, 2*time.Second)
```
![](demo.gif)