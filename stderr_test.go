package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestStdErr(t *testing.T) {
	exp := 1.23659479
	act := StdErr([]float64{2.3, 3.5, 5.7, 7.9})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
