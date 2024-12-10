package util

func GetFrequencyMap[T comparable](right []T) map[T]int {
	m := make(map[T]int)
	for _, v := range right {
		m[v]++
	}
	return m
}
