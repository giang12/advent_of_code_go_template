package day_07

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
func parseInput(input []string) (values []int, nums [][]int) {
	values = make([]int, len(input))
	nums = make([][]int, len(input))
	for i, str := range input {
		tokens := strings.Split(str, ":")
		values[i] = util.Atoi(tokens[0])
		nums[i] = util.Map(strings.Split(strings.Trim(tokens[1], " "), " "), func(item string) int { return util.Atoi(item) })
	}
	return values, nums
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	defer util.Timer()()
	values, nums := parseInput(input)
	total := 0
	for i, v := range values {
		if isGood1(nums[i], v, nums[i][0], 1) {
			total += v
		}
	}
	return strconv.Itoa(total)
}

func isGood1(nums []int, value int, currVal int, currIdx int) bool {
	if currIdx >= len(nums) {
		return currVal == value
	}
	if currVal > value {
		return false
	}
	plus := isGood1(nums, value, currVal+nums[currIdx], currIdx+1)
	mul := isGood1(nums, value, currVal*nums[currIdx], currIdx+1)
	return plus || mul
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	defer util.Timer()()
	values, nums := parseInput(input)
	total := 0
	for i, v := range values {
		if isGood2(nums[i], v, nums[i][0], 1) {
			total += v
		}
	}
	return strconv.Itoa(total)
}
func isGood2(nums []int, value int, currVal int, currIdx int) bool {
	if currIdx >= len(nums) {
		return currVal == value
	}
	if currVal > value {
		return false
	}
	plus := isGood2(nums, value, currVal+nums[currIdx], currIdx+1)
	mul := isGood2(nums, value, currVal*nums[currIdx], currIdx+1)
	concat := isGood2(nums, value, util.Atoi(util.Itoa(currVal)+util.Itoa(nums[currIdx])), currIdx+1)
	return plus || mul || concat
}
