package list

import "github.com/MFQWKMR4/MyFPGo/pkg/hkt"

func Filter[A any](f func(A) bool) func(hkt.K1[List, A]) hkt.K1[List, A] {
	return func(k hkt.K1[List, A]) hkt.K1[List, A] {
		switch k := k.(type) {
		case NonEmptyT[A]:
			values := make([]A, 0)
			for _, v := range k.Values {
				if f(v) {
					values = append(values, v)
				}
			}
			if len(values) == 0 {
				return E[A]()
			}
			return NonEmpty[A](values)
		case ET[A]:
			return E[A]()
		default:
			panic("unreachable")
		}
	}
}

func ForEach[A any](f func(A)) func(hkt.K1[List, A]) {
	return func(k hkt.K1[List, A]) {
		switch k := k.(type) {
		case NonEmptyT[A]:
			for _, v := range k.Values {
				f(v)
			}
		case ET[A]:
			return
		default:
			panic("unreachable")
		}
	}
}
