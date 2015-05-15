package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestMax(t *testing.T) {
	exp := 7.9
	act := Max([]float64{5.7, 2.3, 7.9, 3.5})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
