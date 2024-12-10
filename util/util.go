package util

import (
	"math"
	"strconv"
)

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
