package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestBinDist(t *testing.T) {
	exp := 0.34375
	act := BinDist(0.5, 6, 0, 2)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
