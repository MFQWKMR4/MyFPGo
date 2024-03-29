package list

import "github.com/MFQWKMR4/MyFPGo/pkg/hkt"

type List any

type NonEmptyT[A any] struct {
	Values []A
}

func NonEmpty[A any](values []A) NonEmptyT[A] {
	return NonEmptyT[A]{Values: values}
}

func (NonEmptyT[A]) HKT1(List) {}
func (NonEmptyT[A]) HKT2(A)    {}

type ET[A any] struct{}

func E[A any]() ET[A] {
	return ET[A]{}
}

func (ET[A]) HKT1(List) {}
func (ET[A]) HKT2(A)    {}

func New[A any](list []A) hkt.K1[List, A] {
	if len(list) == 0 {
		return E[A]()
	}
	return NonEmpty[A](list)
}

func Map[A, B any](f func(A) B) func(hkt.K1[List, A]) hkt.K1[List, B] {
	return func(k hkt.K1[List, A]) hkt.K1[List, B] {
		switch k := k.(type) {
		case NonEmptyT[A]:
			values := make([]B, len(k.Values))
			for i, v := range k.Values {
				values[i] = f(v)
			}
			return NonEmpty[B](values)
		case ET[A]:
			return E[B]()
		default:
			panic("unreachable")
		}
	}
}

func Flatten[A any](k hkt.K1[List, hkt.K1[List, A]]) hkt.K1[List, A] {
	switch k := k.(type) {
	case NonEmptyT[hkt.K1[List, A]]:
		values := make([]A, 0)
		for _, v := range k.Values {
			switch v := v.(type) {
			case NonEmptyT[A]:
				values = append(values, v.Values...)
			case ET[A]:
				continue
			default:
				panic("unreachable")
			}
		}
		return NonEmpty[A](values)
	case ET[hkt.K1[List, A]]:
		return E[A]()
	default:
		panic("unreachable")
	}
}

func FlatMap[A, B any](f func(A) hkt.K1[List, B]) func(hkt.K1[List, A]) hkt.K1[List, B] {
	return func(k hkt.K1[List, A]) hkt.K1[List, B] {
		return Flatten[B](Map[A, hkt.K1[List, B]](f)(k))
	}
}

type Functor[A, B any] struct {
	f func(A) B
}

type Monad[A, B any] struct {
	f func(A) hkt.K1[List, B]
}

func F[A, B any](f func(A) B) Functor[A, B] {
	return Functor[A, B]{f: f}
}

func M[A, B any](f func(A) hkt.K1[List, B]) Monad[A, B] {
	return Monad[A, B]{f: f}
}

func (f Functor[A, B]) Map(a hkt.K1[List, A]) hkt.K1[List, B] {
	return Map[A, B](f.f)(a)
}

func (m Monad[A, B]) Map(a hkt.K1[List, A]) hkt.K1[List, B] {
	return FlatMap[A, B](m.f)(a)
}
