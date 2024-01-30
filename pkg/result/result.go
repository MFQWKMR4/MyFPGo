package result

import (
	"github.com/MFQWKMR4/MyFPGo/pkg/hkt"
)

type Result any

type OkT[A any, B error] struct{ Value A }

func Ok[A any, B error](a A) OkT[A, B] {
	return OkT[A, B]{Value: a}
}

func (OkT[A, B]) HKT1(Result) {}
func (OkT[A, B]) HKT2(A)      {}
func (OkT[A, B]) HKT3(B)      {}

type ErrT[A any, B error] struct{ Value error }

func (e ErrT[A, B]) Error() string {
	return e.Value.Error()
}

func Err[A any, B error](b error) ErrT[A, B] {
	return ErrT[A, B]{Value: b}
}

func (ErrT[A, B]) HKT1(Result) {}
func (ErrT[A, B]) HKT2(A)      {}
func (ErrT[A, B]) HKT3(B)      {}

func Map[A any, B error, C any, D error](f func(A) C) func(hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return func(k hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
		switch k := k.(type) {
		case OkT[A, B]:
			return Ok[C, D](f(k.Value))
		case ErrT[A, B]:
			return Err[C, D](k.Value)
		default:
			panic("unreachable")
		}
	}
}

func FlatMap[A any, B error, C any, D error](f func(A) hkt.K2[Result, C, D]) func(hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return func(k hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
		switch k := k.(type) {
		case OkT[A, B]:
			tmp := f(k.Value)
			switch tmp := tmp.(type) {
			case OkT[C, D]:
				return Ok[C, D](tmp.Value)
			case ErrT[C, D]:
				return Err[C, D](tmp.Value)
			default:
				panic("unreachable")
			}
		case ErrT[A, B]:
			return Err[C, D](k.Value)
		default:
			panic("unreachable")
		}
	}
}

type Functor[A any, B error, C any, D error] struct {
	f func(A) C
}

type Monad[A any, B error, C any, D error] struct {
	f func(A) hkt.K2[Result, C, D]
}

func F[A any, B error, C any, D error](f func(A) C) Functor[A, B, C, D] {
	return Functor[A, B, C, D]{f: f}
}

// wrapper of func F
func Fw[A, C any](f func(A) C) Functor[A, error, C, error] {
	return Functor[A, error, C, error]{f: f}
}

func M[A any, B error, C any, D error](f func(A) hkt.K2[Result, C, D]) Monad[A, B, C, D] {
	return Monad[A, B, C, D]{f: f}
}

// wrapper of func M
func Mw[A, C any](f func(A) hkt.K2[Result, C, error]) Monad[A, error, C, error] {
	return Monad[A, error, C, error]{f: f}
}

func (f Functor[A, B, C, D]) Map(a hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return Map[A, B, C, D](f.f)(a)
}

func (m Monad[A, B, C, D]) Map(a hkt.K2[Result, A, B]) hkt.K2[Result, C, D] {
	return FlatMap[A, B, C, D](m.f)(a)
}
