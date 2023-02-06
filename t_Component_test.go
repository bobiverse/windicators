package windicators

import (
	"testing"
	"time"
)

func TestComponent(t *testing.T) {
	c := NewComponent(nil, "", "", nil, time.Second)
	if c != nil {
		t.Fatalf("component should not be made")
	}

	iw, _ := NewIndicatorWindow(1, 1, 1)

	c = NewComponent(iw, "", "", nil, time.Second)
	if c != nil {
		t.Fatalf("component should not be made")
	}

	c = NewComponent(iw, "test", "", nil, time.Second)
	if c != nil {
		t.Fatalf("component should not be made")
	}

	c = NewComponent(iw, "test", "Teting %v", nil, time.Second)
	if c != nil {
		t.Fatalf("component should not be made")
	}

	c = NewComponent(iw, "test", "Teting %v", func(c *Component) any {
		return true
	}, time.Second)

	if c == nil {
		t.Fatalf("component SHOULD BE made")
	}
	if c.ComponentMetrics == nil {
		t.Fatalf("component metrics SHOULD BE made")
	}
	if c.IsValid() {
		t.Fatalf("component metrics SHOULD NOT BE valid yet")
	}

}
