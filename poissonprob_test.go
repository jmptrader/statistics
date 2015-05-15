package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestPoissonProb(t *testing.T) {
	exps := []float64{0.006737947, 0.033689735, 0.084224337, 0.1403739}
	for i := 0; i <= 3; i++ {
		act := PoissonProb(5, i)
		if !calc.Near(exps[i], act) {
			t.Error("Expected", exps[i], "got", act)
		}
	}
}
