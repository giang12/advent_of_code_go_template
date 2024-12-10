package util

func Ternary[T any](condition bool, valueIfTrue, valueIfFalse T) T {
	if condition {
		return valueIfTrue
	}
	return valueIfFalse
}
