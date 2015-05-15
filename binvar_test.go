package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestBinVar(t *testing.T) {
	exp := 25.0
	act := BinVar(100, 0.5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
