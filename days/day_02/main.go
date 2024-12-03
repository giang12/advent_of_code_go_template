package day_02

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

func isReportSafe(levels []string) (bool, int) {
	var inc int = util.Atoi(levels[1]) - util.Atoi(levels[0])

	var i int = 1
	if inc == 0 {
		return false, i
	}

	for ; i < len(levels); i++ {
		var diff int = util.Atoi(levels[i]) - util.Atoi(levels[i-1])
		var abs int = util.Abs(diff)

		//bounds
		if abs < 1 || abs > 3 {
			return false, i
		}
		//different direction
		if (diff < 0 && inc > 0) || (diff > 0 && inc < 0) {
			return false, i
		}
	}
	return true, -1
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	var ans int = 0
	for _, report := range input {
		isSafe, _ := isReportSafe(strings.Fields(report))
		if isSafe {
			ans += 1
		}
	}
	return strconv.Itoa(ans)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	var ans int = 0
	for _, report := range input {
		var cur []string = strings.Fields(report)
		isSafe, _ := isReportSafe(cur)
		if !isSafe {
			for i := 0; i < len(cur); i++ {
				test_report_input := util.RemoveIndex(cur, i)
				isSafe, _ = isReportSafe(test_report_input)
				if isSafe {
					break
				}
			}
		}
		if isSafe {
			ans += 1
		}
	}
	return strconv.Itoa(ans)
}
