package util

func GetFrequencyMap(right []int) map[int]int {
	m := make(map[int]int)
	for _, v := range right {
		m[v]++
	}
	return m
}
