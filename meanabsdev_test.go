package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestMeanAbsDev(t *testing.T) {
	exp := 1.95
	act := MeanAbsDev([]float64{2.3, 3.5, 5.7, 7.9})
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
