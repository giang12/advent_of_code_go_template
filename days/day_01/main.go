package day_01

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"advent_of_code_go_template/util"
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

func parseInput(input []string) ([]int, []int) {
	var lefts []int
	var rights []int
	for i := 0; i < len(input); i++ {
		parts := strings.Fields(input[i])
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		lefts = append(lefts, left)
		rights = append(rights, right)
	}
	return lefts, rights
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	var lefts, rights []int = parseInput(input)

	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})
	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})
	var total int = 0

	for i := 0; i < len(lefts); i++ {
		total += int(math.Abs(float64(lefts[i] - rights[i])))
	}
	//fmt.Printf("leftSet: %v\n", leftSet)
	//fmt.Printf("rightSet: %v\n", rightSet)
	return strconv.Itoa(total)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	var lefts, rights []int = parseInput(input)
	var freqs map[int]int = util.GetFrequencyMap(rights)
	var ans int = 0
	for _, v := range lefts {
		ans += (v * freqs[v])
	}
	return strconv.Itoa(ans)
}
