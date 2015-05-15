package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestVarPop(t *testing.T) {
	exp := 4.5875
	act := VarPop([]float64{2.3, 3.5, 5.7, 7.9})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
