package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestSumSqrDev(t *testing.T) {
	exp := 18.35
	act := SumSqrDev([]float64{2.3, 3.5, 5.7, 7.9})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
