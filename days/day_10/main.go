package day_10

import (
	"advent_of_code_go_template/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
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

func parseInput(input []string) (grid [][]int, startNodes [][2]int) {

	grid = make([][]int, len(input))
	startNodes = make([][2]int, 0)
	for i, line := range input {
		grid[i] = util.Map(strings.Split(line, ""), func(i string) int {
			if i == "." {
				return -1
			}
			return util.Atoi(i)
		})
		for j, val := range grid[i] {
			if val == 0 {
				startNodes = append(startNodes, [2]int{i, j})
			}
		}
	}
	return grid, startNodes
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	grid, startNodes := parseInput(input)

	total := 0
	for _, inital := range startNodes {
		total += count(grid, inital[0], inital[1], true)
	}
	return strconv.Itoa(total)
}

var DIRS = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func count(grid [][]int, r int, c int, uniqPath bool) int {
	mem := make([][]bool, len(grid))
	for i := range mem {
		mem[i] = make([]bool, len(grid[0]))
	}
	stack := arraystack.New() // empty
	stack.Push([2]int{r, c})  // 1
	mem[r][c] = true
	total := 0
	for !stack.Empty() {
		item, _ := stack.Pop()
		curr := item.([2]int)
		if grid[curr[0]][curr[1]] == 9 {
			total++
			continue
		}
		for _, d := range DIRS {
			nextR := curr[0] + d[0]
			nextC := curr[1] + d[1]
			if nextR < 0 || nextC < 0 || nextR >= len(grid) || nextC >= len(grid[0]) {
				continue
			}
			if uniqPath && mem[nextR][nextC] {
				continue
			}
			if grid[nextR][nextC]-grid[curr[0]][curr[1]] == 1 {
				mem[nextR][nextC] = true
				stack.Push([2]int{nextR, nextC})
			}
		}
	}
	return total
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	grid, startNodes := parseInput(input)

	total := 0
	for _, inital := range startNodes {
		total += count(grid, inital[0], inital[1], false)
	}
	return strconv.Itoa(total)
}
