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
