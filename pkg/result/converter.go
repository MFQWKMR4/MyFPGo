package result

import "github.com/MFQWKMR4/MyFPGo/pkg/hkt"

func From[A any, B error](a *A, b *B) hkt.K2[Result, A, B] {

	if b != nil {
		return Err[A, B](*b)
	}
	return Ok[A, B](*a)
}

func To[A any, B error](k hkt.K2[Result, A, B]) (*A, *B) {
	switch k := k.(type) {
	case OkT[A, B]:
		return &k.Value, nil
	case ErrT[A, B]:
		err, ok := k.Value.(B) // TODO: check type
		if !ok {
			panic("unexpected error type")
		}
		return nil, &err
	default:
		panic("unreachable")
	}
}
