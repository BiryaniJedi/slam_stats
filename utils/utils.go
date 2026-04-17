package utils

func Map[T, E any](xs []T, fn func(T) E) []E {
	ret := []E{}
	for _, x := range xs {
		ret = append(ret, fn(x))
	}
	return ret
}

func MapSafe[T, E any](xs []T, fn func(T) (E, error)) ([]E, []error) {
	ret := []E{}
	errors := []error{}
	for _, x := range xs {
		eVal, err := fn(x)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		ret = append(ret, eVal)
	}
	return ret, errors
}
