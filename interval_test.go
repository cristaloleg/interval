package interval

import "testing"

func TestInterval(t *testing.T) {
	a := NewIntr(1, 2)
	b := NewIntr(5, 7)
	x := NewIntr(2, 3)

	if value := a.Mul(x).Add(b); value.x != 7 || value.y != 13 {
		t.Errorf("wut")
	}
}
