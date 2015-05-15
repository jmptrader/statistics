package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestPoissonVar(t *testing.T) {
	exp := 5.0
	act := PoissonVar(5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
