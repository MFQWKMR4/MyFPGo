package hkt

// K1 is a higher-kinded type with one type parameter.
type K1[F, A any] interface {
	HKT1(F)
	HKT2(A)
}

// K2 is a higher-kinded type with two type parameters.
type K2[F, A, B any] interface {
	HKT1(F)
	HKT2(A)
	HKT3(B)
}

type Functor[F, A, B any] interface {
	Map(K1[F, A]) K1[F, B]
}

type Foldable[F, A, B any] interface {
	// from f: func(A, B) A
	FoldLeft(A, K1[F, B]) A
}

type Functor2[F, A, B, C, D any] interface {
	Map(K2[F, A, B]) K2[F, C, D]
}

func Chain2[F, A, B, C any](f Functor[F, A, B], g Functor[F, B, C]) func(K1[F, A]) K1[F, C] {
	return func(k K1[F, A]) K1[F, C] {
		return g.Map(f.Map(k))
	}
}

func Chain3[F, A, B, C, D any](f Functor[F, A, B], g Functor[F, B, C], h Functor[F, C, D]) func(K1[F, A]) K1[F, D] {
	return func(k K1[F, A]) K1[F, D] {
		return h.Map(g.Map(f.Map(k)))
	}
}

func Chain22[F, A, B, C, D, E, G any](f Functor2[F, A, B, C, D], g Functor2[F, C, D, E, G]) func(K2[F, A, B]) K2[F, E, G] {
	return func(k K2[F, A, B]) K2[F, E, G] {
		return g.Map(f.Map(k))
	}
}

func Switch[F, A, B, C, D any](caseA func(A) K2[F, C, D], caseB func(B) K2[F, C, D]) func(K2[F, A, B]) K2[F, C, D] {
	return func(k K2[F, A, B]) K2[F, C, D] {
		switch k := k.(type) {
		case A:
			return caseA(k)
		case B:
			return caseB(k)
		default:
			panic("unreachable")
		}
	}
}
