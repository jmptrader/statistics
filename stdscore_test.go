package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestStdScore(t *testing.T) {
	exps := []float64{-1.0, 0.0, 1.0, 3.5}
	vals := []float64{4.0, 5.0, 6.0, 8.5}
	mean := 5.0
	stdDev := 1.0
	for i := 0; i < len(exps); i++ {
		act := StdScore(vals[i], mean, stdDev)
		if !calc.Near(exps[i], act) {
			t.Error("Expected", exps[i], "got", act)
		}
	}
}
