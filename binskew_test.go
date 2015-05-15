package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestBinSkew(t *testing.T) {
	exp := 0.0
	act := BinSkew(100, 0.5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
