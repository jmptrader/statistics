package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestBinMean(t *testing.T) {
	exp := 50.0
	act := BinMean(100, 0.5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
