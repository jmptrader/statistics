package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestPoissonSkew(t *testing.T) {
	exp := 0.4472136
	act := PoissonSkew(5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
