package util

func RemoveIndex[T any](cur []T, i int) []T {
	newArr := make([]T, len(cur)-1)
	copy(newArr, cur[:i])
	copy(newArr[i:], cur[i+1:])
	return newArr
}
