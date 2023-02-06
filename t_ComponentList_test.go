package windicators

import (
	"testing"
	"time"
)

func TestComponentList(t *testing.T) {
	iw, _ := NewIndicatorWindow(1, 1, 1)

	c1 := NewComponent(iw, "test", "Teting %v", func(c *Component) any {
		return true
	}, time.Second)
	NewComponent(iw, "test2", "Teting %v", func(c *Component) any {
		return true
	}, time.Second)
	NewComponent(iw, "test3", "Teting %v", func(c *Component) any {
		return true
	}, time.Second)

	// all must be invisible by default
	if iw.components.IsAllHidden() == false {
		t.Fatalf("All must be invisible for now")
	}

	if len(iw.components.FilterVisible()) != 0 {
		t.Fatalf("All must be invisible for now")
	}

	if iw.components.FindByID("test") == nil {
		t.Fatalf("one shoud be found")
	}

	if iw.components.FindByID("xxx") != nil {
		t.Fatalf("this should not be found")
	}

	c1.IsVisible = true

	// all must be invisible by default
	if iw.components.IsAllHidden() == true {
		t.Fatalf("One  must be visible")
	}

	if len(iw.components.FilterVisible()) != 1 {
		t.Fatalf("one must be visible")
	}
}
