package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestStdDev(t *testing.T) {
	exp := 2.47318957
	act := StdDev([]float64{2.3, 3.5, 5.7, 7.9})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
