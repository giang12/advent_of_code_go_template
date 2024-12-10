package day_08

import (
	"advent_of_code_go_template/util"
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

func parseInput(input []string) (grid [][]string, grid2 [][]string, stations map[string][][2]int) {
	grid = make([][]string, len(input))
	grid2 = make([][]string, len(input))
	stations = make(map[string][][2]int)
	for i, line := range input {
		grid[i] = strings.Split(strings.Trim(line, ""), "")
		grid2[i] = make([]string, len(grid[i]))
		for j, char := range grid[i] {
			if char != "." {
				stations[char] = append(stations[char], [2]int{i, j})
			}
			grid2[i][j] = "."
		}
	}
	return grid, grid2, stations
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	defer util.Timer()()
	grid, grid2, stations := parseInput(input)
	//total := 0
	for k, v := range stations {
		mark(grid, grid2, v, k)
	}
	hashtag := 0
	for i := 0; i < len(grid2); i++ {
		//fmt.Println(grid2[i])
		for j := 0; j < len(grid2[0]); j++ {
			if grid2[i][j] == "#" {
				hashtag++
			}
		}
	}
	//fmt.Printf("%v vs %v", hashtag, total)
	return strconv.Itoa(hashtag)
}
func mark(grid [][]string, grid2 [][]string, locations [][2]int, symbol string) int {
	total := 0
	for i := 0; i < len(locations); i++ {
		for j := 0; j < len(locations); j++ {
			total += markPair(grid, grid2, locations[i], locations[j], symbol)
		}
	}
	return total
}
func markPair(grid [][]string, grid2 [][]string, A [2]int, B [2]int, symbol string) int {
	rise := A[1] - B[1]
	run := A[0] - B[0]
	if rise == 0 && run == 0 {
		return 0
	}
	p1 := [2]int{A[0] + run, A[1] + rise}
	p2 := [2]int{A[0] - run, A[1] - rise}
	p3 := [2]int{B[0] + run, B[1] + rise}
	p4 := [2]int{B[0] - run, B[1] - rise}
	return markPoint(grid, grid2, p1, symbol) + markPoint(grid, grid2, p2, symbol) + markPoint(grid, grid2, p3, symbol) + markPoint(grid, grid2, p4, symbol)
}
func markPoint(grid [][]string, grid2 [][]string, A [2]int, symbol string) int {
	if A[0] < 0 || A[1] < 0 || A[0] >= len(grid) || A[1] >= len(grid[0]) {
		return 0
	}
	if grid2[A[0]][A[1]] == "." && grid[A[0]][A[1]] != symbol {
		grid2[A[0]][A[1]] = "#"
	}
	return 1
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	defer util.Timer()()
	grid, grid2, stations := parseInput(input)
	for k, v := range stations {
		mark2(grid, grid2, v, k)
	}
	hashtag := 0
	for i := 0; i < len(grid2); i++ {
		//fmt.Println(grid2[i])
		for j := 0; j < len(grid2[0]); j++ {
			if grid2[i][j] == "#" {
				hashtag++
			}
		}
	}
	//fmt.Printf("%v vs %v", hashtag, total)
	return strconv.Itoa(hashtag)
}
func mark2(grid [][]string, grid2 [][]string, locations [][2]int, symbol string) int {
	total := 0
	for i := 0; i < len(locations); i++ {
		for j := 0; j < len(locations); j++ {
			total += markPair2(grid, grid2, locations[i], locations[j], symbol)
		}
	}
	return total
}
func markPair2(grid [][]string, grid2 [][]string, A [2]int, B [2]int, symbol string) int {
	rise := A[1] - B[1]
	run := A[0] - B[0]
	if rise == 0 && run == 0 {
		return 0
	}
	p1 := [2]int{A[0] + run, A[1] + rise}
	p2 := [2]int{A[0] - run, A[1] - rise}
	for markPoint2(grid, grid2, p1, symbol) > 0 {
		p1 = [2]int{p1[0] + run, p1[1] + rise}
	}
	for markPoint2(grid, grid2, p2, symbol) > 0 {
		p2 = [2]int{p2[0] - run, p2[1] - rise}
	}
	return 0
}
func markPoint2(grid [][]string, grid2 [][]string, A [2]int, symbol string) int {
	if A[0] < 0 || A[1] < 0 || A[0] >= len(grid) || A[1] >= len(grid[0]) {
		return 0
	}
	if grid2[A[0]][A[1]] == "." || grid[A[0]][A[1]] == symbol {
		grid2[A[0]][A[1]] = "#"
	}
	return 1
}
