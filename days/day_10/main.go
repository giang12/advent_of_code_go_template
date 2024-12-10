package day_10

import (
	"advent_of_code_go_template/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/v2/stacks/arraystack"
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
	defer util.Timer()()
	grid, startNodes := parseInput(input)

	total := util.Reduce(startNodes, func(currentVal int, currentItem [2]int) int {
		return currentVal + count(grid, currentItem, false)
	}, 0)

	return strconv.Itoa(total)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	defer util.Timer()()

	grid, startNodes := parseInput(input)

	total := util.Reduce(startNodes, func(currentVal int, currentItem [2]int) int {
		return currentVal + count(grid, currentItem, false)
	}, 0)

	return strconv.Itoa(total)
}

var DIRS = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func count(grid [][]int, start [2]int, uniqPath bool) int {
	mem := make([][]bool, len(grid))
	for i := range mem {
		mem[i] = make([]bool, len(grid[0]))
	}
	stack := *arraystack.New[[2]int]() // empty
	visit(start, mem, stack)
	total := 0
	for !stack.Empty() {
		currLocation, _ := stack.Pop()
		total += util.Ternary(grid[currLocation[0]][currLocation[1]] == 9, 1, 0)
		explore(grid, currLocation, uniqPath, mem, stack)
	}
	return total
}

func explore(grid [][]int, currLocation [2]int, uniqPath bool, mem [][]bool, stack arraystack.Stack[[2]int]) {
	for _, d := range DIRS {
		nextR := currLocation[0] + d[0]
		nextC := currLocation[1] + d[1]
		if nextR < 0 || nextC < 0 || nextR >= len(grid) || nextC >= len(grid[0]) {
			continue
		}
		if uniqPath && mem[nextR][nextC] {
			continue
		}
		if grid[nextR][nextC]-grid[currLocation[0]][currLocation[1]] == 1 {
			visit([2]int{nextR, nextC}, mem, stack)
		}
	}
}
func visit(location [2]int, mem [][]bool, stack arraystack.Stack[[2]int]) {
	mem[location[0]][location[1]] = true
	stack.Push(location) // 1
}
