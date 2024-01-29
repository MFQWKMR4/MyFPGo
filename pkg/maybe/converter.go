package maybe

import "github.com/MFQWKMR4/MyFPGo/pkg/hkt"

func From[A any](a *A) hkt.K1[Maybe, A] {
	if a == nil {
		return None[A]()
	}
	return Some[A](*a)
}

func To[A any](k hkt.K1[Maybe, A]) *A {
	switch k := k.(type) {
	case SomeT[A]:
		return &k.Value
	case NoneT[A]:
		return nil
	default:
		panic("unreachable")
	}
}
