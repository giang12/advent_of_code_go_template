package day_03

import (
	"fmt"
	"regexp"
	"strconv"
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
func Part1(inputs []string) string {
	total := 0
	for _, input := range inputs {
		for _, instr := range getInstructions(input, `mul\(\d+,\d+\)`) {

			total += mul(instr)
		}
	}
	return strconv.Itoa(total)
}

// Part2 solves the second part of the exercise
func Part2(inputs []string) string {
	total := 0
	do := true
	for _, input := range inputs {
		for _, instr := range getInstructions(input, `mul\(\d+,\d+\)|do\(\)|don't\(\)`) {
			if instr == "do()" {
				do = true
			} else if instr == "don't()" {
				do = false
			} else if do {
				total += mul(instr)
			}
		}
	}
	return strconv.Itoa(total)
}

func getInstructions(input string, reg string) []string {
	pattern := regexp.MustCompile(reg)
	return pattern.FindAllString(input, -1)
}
func mul(mul string) int {
	pattern := regexp.MustCompile(`\d+`)
	var numbers []int
	for _, s := range pattern.FindAllString(mul, -1) {
		num, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			continue
		}

		numbers = append(numbers, num)
	}
	return numbers[0] * numbers[1]
}
