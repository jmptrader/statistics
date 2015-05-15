package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestMin(t *testing.T) {
	exp := 2.3
	act := Min([]float64{5.7, 2.3, 7.9, 3.5})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
