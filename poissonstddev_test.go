package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestPoissonStdDev(t *testing.T) {
	exp := 2.236068
	act := PoissonStdDev(5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
