package services

func Map[IN any, OUT any](array []IN, fn func(t IN) OUT) []OUT {
	result := make([]OUT, len(array))
	for i, item := range array {
		result[i] = fn(item)
	}
	return result
}

func Contains[T int | string](array []T, value T) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}
