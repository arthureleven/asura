package prototyping

func Map[T any](slice []T, callback func(T) T) []T {
	s := make([]T, len(slice))

	for i, elem := range slice {
		s[i] = callback(elem)
	}

	return s
}
