package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestChebyshev(t *testing.T) {
	exps := []float64{0.0, 0.0, 0.75, 0.88888889, 0.9375, 0.96}
	for i := 0; i < 6; i++ {
		act := Chebyshev(float64(i))
		if !calc.Near(exps[i], act) {
			t.Error("Expected", exps[i], "got", act)
		}
	}
}
