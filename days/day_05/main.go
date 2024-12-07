package day_05

import (
	"advent_of_code_go_template/util"
	"fmt"
	"sort"
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
func parseInput(input []string) (requirements map[int]map[int]bool, updates []string) {
	r := make(map[int]map[int]bool)
	var currI int
	for index, i := range input {
		currI = index
		if len(i) == 0 {
			break
		}
		numStrings := strings.Split(i, "|")
		if r[util.Atoi(numStrings[0])] == nil {
			r[util.Atoi(numStrings[0])] = make(map[int]bool)
		}
		r[util.Atoi(numStrings[0])][util.Atoi(numStrings[1])] = true
	}

	currI++
	u := input[currI:]
	return r, u
}
func getValidMid(reqs map[int]map[int]bool, update string) int {
	nums := util.Map(strings.Split(update, ","), func(item string) int { return util.Atoi(item) })

	if isValidUpdate(reqs, nums) {
		return nums[len(nums)/2]
	}
	return 0
}
func isValidUpdate(reqs map[int]map[int]bool, nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if reqs[nums[j]][nums[i]] {
				return false
			}
		}
	}
	return true
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	defer util.Timer()()

	reqs, updates := parseInput(input)
	total := 0

	for _, u := range updates {
		total += getValidMid(reqs, u)
	}
	return strconv.Itoa(total)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	defer util.Timer()()

	reqs, updates := parseInput(input)
	total := 0

	for _, u := range updates {
		if getValidMid(reqs, u) == 0 {
			//fixUpdate(reqs, u)
			total += getValidMid(reqs, fixUpdate(reqs, u))
		}
	}
	return strconv.Itoa(total)
}
func fixUpdate(reqs map[int]map[int]bool, update string) string {
	nums := util.Map(strings.Split(update, ","), func(item string) int { return util.Atoi(item) })

	//dont do bubblesort kids lmao..
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {

	// 		if reqs[nums[j]][nums[i]] {
	// 			swap(reqs, nums, i, j)
	// 		}
	// 	}
	// }
	sort.Slice(nums, func(i, j int) bool {
		return !reqs[nums[j]][nums[i]]
	})

	numsString := util.Map(nums, func(item int) string { return util.Itoa(item) })
	return strings.Join(numsString, ",")
}
func swap(reqs map[int]map[int]bool, nums []int, a int, b int) {
	tmp := nums[b]
	nums[b] = nums[a]
	nums[a] = tmp
}
