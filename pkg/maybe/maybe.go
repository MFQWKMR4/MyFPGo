package maybe

import "github.com/MFQWKMR4/MyFPGo/pkg/hkt"

type Maybe any

type SomeT[A any] struct{ Value A }

func Some[A any](a A) SomeT[A] {
	return SomeT[A]{Value: a}
}

func (SomeT[A]) HKT1(Maybe) {}
func (SomeT[A]) HKT2(A)     {}

type NoneT[A any] struct{}

func None[A any]() NoneT[A] {
	return NoneT[A]{}
}

func (NoneT[A]) HKT1(Maybe) {}
func (NoneT[A]) HKT2(A)     {}

func Map[A, B any](f func(A) B) func(hkt.K1[Maybe, A]) hkt.K1[Maybe, B] {
	return func(k hkt.K1[Maybe, A]) hkt.K1[Maybe, B] {
		switch k := k.(type) {
		case SomeT[A]:
			return Some[B](f(k.Value))
		case NoneT[A]:
			return None[B]()
		default:
			panic("unreachable")
		}
	}
}

func FlatMap[A, B any](f func(A) hkt.K1[Maybe, B]) func(hkt.K1[Maybe, A]) hkt.K1[Maybe, B] {
	return func(k hkt.K1[Maybe, A]) hkt.K1[Maybe, B] {
		switch k := k.(type) {
		case SomeT[A]:
			return f(k.Value)
		case NoneT[A]:
			return None[B]()
		default:
			panic("unreachable")
		}
	}
}

type Functor[A, B any] struct {
	f func(A) B
}

type Monad[A, B any] struct {
	f func(A) hkt.K1[Maybe, B]
}

func F[A, B any](f func(A) B) Functor[A, B] {
	return Functor[A, B]{f: f}
}

func M[A, B any](f func(A) hkt.K1[Maybe, B]) Monad[A, B] {
	return Monad[A, B]{f: f}
}

func (f Functor[A, B]) Map(a hkt.K1[Maybe, A]) hkt.K1[Maybe, B] {
	return Map[A, B](f.f)(a)
}

// NOTE: Monadといいつつ、Mapだけ実装
func (m Monad[A, B]) Map(a hkt.K1[Maybe, A]) hkt.K1[Maybe, B] {
	return FlatMap[A, B](m.f)(a)
}
