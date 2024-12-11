package day_11

import (
	"advent_of_code_go_template/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/v2/maps/hashmap"
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

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	defer util.Timer()()
	return solve(strings.Split(input[0], " "), 25)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	defer util.Timer()()
	return solve(strings.Split(input[0], " "), 75)
}

func solve(arr []string, iters int) string {
	total := 0
	m := hashmap.New[string, int]() // empty
	for _, num := range arr {
		total += count(num, iters, m)
	}
	return strconv.Itoa(total)
}

// return number of splits after iters
func count(val string, iters int, mem *hashmap.Map[string, int]) int {
	if iters == 0 {
		return 1
	}
	key := fmt.Sprintf("%v@%v", val, iters)
	cacheVal, found := mem.Get(key)
	if found {
		// fmt.Printf("%v -> %v\n", key, cacheVal)
		return cacheVal
	}

	cnt := 0

	if val == "0" {
		cnt = count("1", iters-1, mem)
	} else if len(val)%2 == 0 {
		mid := len(val) / 2
		left := util.Itoa(util.Atoi(val[:mid]))
		right := util.Itoa(util.Atoi(val[mid:]))
		cnt = count(left, iters-1, mem) + count(right, iters-1, mem)
	} else {
		cnt = count(util.Itoa(util.Atoi(val)*2024), iters-1, mem)
	}

	mem.Put(key, cnt)
	return cnt
}
