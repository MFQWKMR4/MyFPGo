package result

import (
	"github.com/MFQWKMR4/MyFPGo/pkg/hkt"
)

type Result any

type OkT[A, B any] struct{ Value A }

func Ok[A, B any](a A) OkT[A, B] {
	return OkT[A, B]{Value: a}
}

func (OkT[A, B]) HKT1(Result) {}
func (OkT[A, B]) HKT2(A)      {}
func (OkT[A, B]) HKT3(B)      {}

type ErrT[A, B any] struct{ Value B }

func Err[A, B any]() ErrT[A, B] {
	return ErrT[A, B]{}
}

func (ErrT[A, B]) HKT1(Result) {}
func (ErrT[A, B]) HKT2(A)      {}
func (ErrT[A, B]) HKT3(B)      {}

func Map[A, B, C, D any](f func(A) C) func(hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return func(k hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
		switch k := k.(type) {
		case OkT[A, B]:
			return Ok[C, D](f(k.Value))
		case ErrT[A, B]:
			return Err[C, D]()
		default:
			panic("unreachable")
		}
	}
}

func FlatMap[A, B, C, D any](f func(A) hkt.K2[Result, C, D]) func(hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return func(k hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
		switch k := k.(type) {
		case OkT[A, B]:
			tmp := f(k.Value)
			switch tmp := tmp.(type) {
			case OkT[C, D]:
				return Ok[C, D](tmp.Value)
			case ErrT[C, D]:
				return Err[C, D]()
			default:
				panic("unreachable")
			}
		case ErrT[A, B]:
			return Err[C, D]()
		default:
			panic("unreachable")
		}
	}
}

type Functor[A, B, C, D any] struct {
	f func(A) C
}

type Monad[A, B, C, D any] struct {
	f func(A) hkt.K2[Result, C, D]
}

func F[A, B, C, D any](f func(A) C) Functor[A, B, C, D] {
	return Functor[A, B, C, D]{f: f}
}

func M[A, B, C, D any](f func(A) hkt.K2[Result, C, D]) Monad[A, B, C, D] {
	return Monad[A, B, C, D]{f: f}
}

func (f Functor[A, B, C, D]) Map(a hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return Map[A, B, C, D](f.f)(a)
}

func (m Monad[A, B, C, D]) Map(a hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return FlatMap[A, B](m.f)(a)
}
