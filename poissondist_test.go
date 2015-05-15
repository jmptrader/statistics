package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestPoissonDist(t *testing.T) {
	exp := 0.26502592
	act := PoissonDist(5, 0, 3)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
