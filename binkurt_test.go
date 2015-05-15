package stat

import (
	"testing"

	"github.com/BluntSporks/calculation"
)

func TestBinKurt(t *testing.T) {
	exp := 2.98
	act := BinKurt(100, 0.5)
	if !calc.Near(exp, act) {
		t.Error("Expected", exp, "got", act)
	}
}
