package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestBinStdDev(t *testing.T) {
	exp := 5.0
	act := BinStdDev(100, 0.5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
