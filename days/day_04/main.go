package day_04

import (
	"fmt"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}
func parseInput(input []string) [][]string {
	arg := make([][]string, len(input))
	for i, val := range input {
		arg[i] = strings.Split(val, "")
	}
	return arg
}

var ref = []string{"X", "M", "A", "S"}
var dir = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1}}

func check1(arg [][]string, x int, y int) int {
	cnt := 0
	for _, d := range dir {
		found := true
		currX := x
		currY := y
		for _, char := range ref {
			if currX < 0 || currY < 0 || currX >= len(arg) || currY >= len(arg[0]) {
				found = false
				break
			}
			if arg[currX][currY] != char {
				found = false
				break
			}
			currX += d[0]
			currY += d[1]
		}
		if found {
			cnt++
		}
	}
	return cnt
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	arg := parseInput(input)
	N := len(arg)
	M := len(arg[0])
	total := 0
	for x := 0; x < N; x++ {
		for y := 0; y < M; y++ {
			if arg[x][y] == "X" {
				//fmt.Printf("check %v, %v\n", x, y)
				total += check1(arg, x, y)
			}
		}
	}
	return strconv.Itoa(total)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	arg := parseInput(input)
	N := len(arg)
	M := len(arg[0])
	total := 0
	for x := 0; x < N; x++ {
		for y := 0; y < M; y++ {
			if arg[x][y] == "A" && check2(arg, x, y) {
				//fmt.Printf("check %v, %v\n", x, y)
				total += 1
			}
		}
	}
	return strconv.Itoa(total)
}

func check2(arg [][]string, x int, y int) bool {

	return checkPair(arg, x-1, y-1, x+1, y+1) && checkPair(arg, x+1, y-1, x-1, y+1)
}
func checkPair(arg [][]string, x int, y int, xx int, yy int) bool {
	if x < 0 || y < 0 || x >= len(arg) || y >= len(arg[0]) {
		return false
	}
	if xx < 0 || yy < 0 || xx >= len(arg) || yy >= len(arg[0]) {
		return false
	}
	return (arg[x][y] == "S" && arg[xx][yy] == "M") || (arg[x][y] == "M" && arg[xx][yy] == "S")
}
