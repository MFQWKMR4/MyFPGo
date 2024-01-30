package maybe

import (
	"testing"

	"github.com/MFQWKMR4/MyFPGo/pkg/hkt"
)

func MaybeEqTest(t *testing.T) {
	var a, b, c hkt.K1[Maybe, int] = Some(1), Some(1), Some(2)
	if !(a == b) {
		t.Errorf("expected %v == %v", a, b)
	}
	if a == c {
		t.Errorf("expected %v != %v", a, c)
	}
}
