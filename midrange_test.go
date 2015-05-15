package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestMidrange(t *testing.T) {
	exp1 := 5.1
	act1 := Midrange([]float64{7.9, 2.3, 5.7})
	if !calc.Near(exp1, act1) {
		t.Error("Expected", exp1, "got", act1)
	}
	exp2 := 5.7
	act2 := Midrange([]float64{7.9, 2.3, 5.2, 9.1})
	if !calc.Near(exp2, act2) {
		t.Error("Expected", exp2, "got", act2)
	}
}
