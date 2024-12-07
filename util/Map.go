package util

/**
util.Map([], func(item string) int { return util.Atoi(item) })
**/
func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
