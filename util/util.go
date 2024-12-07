package util

import (
	"math"
	"strconv"
)

func GetFrequencyMap(right []int) map[int]int {
	m := make(map[int]int)
	for _, v := range right {
		m[v]++
	}
	return m
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func Itoa(s int) string {
	return strconv.Itoa(s)
}
func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func RemoveIndex(cur []string, i int) []string {
	newArr := make([]string, len(cur)-1)
	copy(newArr, cur[:i])
	copy(newArr[i:], cur[i+1:])
	return newArr
}
