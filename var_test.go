package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestVar(t *testing.T) {
	exp := 6.11666667
	act := Var([]float64{2.3, 3.5, 5.7, 7.9})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
