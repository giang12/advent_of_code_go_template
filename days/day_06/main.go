package day_06

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

var dirsDelta = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func parseInput(input []string) (grids [][]string, r int, c int, dir int, numsOb int) {
	grids = make([][]string, len(input))
	numsOb = 0
	for i, line := range input {
		grids[i] = strings.Split(line, "")
		for j, char := range grids[i] {
			if char == "^" {
				r = i
				c = j
				dir = 0
			} else if char == ">" {
				r = i
				c = j
				dir = 1
			} else if char == "v" {
				r = i
				c = j
				dir = 2
			} else if char == "<" {
				r = i
				c = j
				dir = 3
			} else if isObstacle(grids, i, j) {
				numsOb++
			}
		}
	}

	return grids, r, c, dir, numsOb
}
func isObstacle(grids [][]string, r int, c int) bool {
	return grids[r][c] == "#"
}
func visited(grids [][]string, r int, c int) bool {
	return grids[r][c] == "X"
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	defer util.Timer()()
	grids, r, c, dir, _ := parseInput(input)
	return strconv.Itoa(solve1(grids, r, c, dir))
}
func solve1(grids [][]string, r int, c int, dir int) int {
	cnt := 0
	for r >= 0 && c >= 0 && r < len(grids) && c < len(grids[0]) {
		if !visited(grids, r, c) {
			cnt++
			grids[r][c] = "X"
		}
		nextR := r + dirsDelta[dir][0]
		nextC := c + dirsDelta[dir][1]
		if !(nextR >= 0 && nextC >= 0 && nextR < len(grids) && nextC < len(grids[0])) {
			break
		}
		if isObstacle(grids, nextR, nextC) {
			dir = (dir + 1) % len(dirsDelta)
		} else {
			r = nextR
			c = nextC
		}
	}
	return cnt
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	defer util.Timer()()
	grids, r, c, dir, numsOb := parseInput(input)
	ans := 0
	for i := range grids {
		for j := range grids[0] {
			if i == r && j == c {
				continue //starting
			}
			if grids[i][j] == "." {
				grids[i][j] = "#"
				if solve2(grids, r, c, dir, numsOb+1) {
					ans++
				}
				grids[i][j] = "."
			}
		}
	}
	// grids[6][3] = "#"
	// if solve2(grids, r, c, dir, numsOb+1) {
	// 	ans++
	// }
	return strconv.Itoa(ans)
}
func solve2(grids [][]string, r int, c int, dir int, numsOb int) bool {
	LEN := len(grids) * len(grids[0])
	cnt := 0
	//visited := [1000000][100000]bool{}
	for cnt+numsOb <= LEN {
		//if !visited[r][c] {
		cnt++
		//	visited[r][c] = true
		//}
		nextR := r + dirsDelta[dir][0]
		nextC := c + dirsDelta[dir][1]
		if !(nextR >= 0 && nextC >= 0 && nextR < len(grids) && nextC < len(grids[0])) {
			break
		}
		if isObstacle(grids, nextR, nextC) {
			dir = (dir + 1) % len(dirsDelta)
		} else {
			r = nextR
			c = nextC
		}
	}
	return cnt+numsOb > LEN
}
