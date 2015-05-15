package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestBinProb(t *testing.T) {
	exp := 0.234375
	act := BinProb(0.5, 6, 2)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
