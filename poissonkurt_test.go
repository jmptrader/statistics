package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestPoissonKurt(t *testing.T) {
	exp := 3.2
	act := PoissonKurt(5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
