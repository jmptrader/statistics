package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestPoissonMean(t *testing.T) {
	exp := 5.0
	act := PoissonMean(5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
