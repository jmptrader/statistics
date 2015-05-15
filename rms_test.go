package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestRMS(t *testing.T) {
	exp := 5.3018865
	act := RMS([]float64{2.3, 3.5, 5.7, 7.9})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
