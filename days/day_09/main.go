package day_09

import (
	"advent_of_code_go_template/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/v2/sets/treeset"
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
func parseInput(input []string) (arr []int, available_spaces treeset.Set[Space], used_spaces []Space, maxID int) {
	freeSpace := false
	maxID = -1
	arr = make([]int, 0)
	available_spaces = *treeset.NewWith[Space](bySpace) // empty (keys are of type int)
	used_spaces = make([]Space, 0)
	for _, s := range input {
		cnt := util.Atoi(s)

		if freeSpace {
			if cnt > 0 {
				available_spaces.Add(Space{-1, len(arr), len(arr) + cnt - 1})
				for cnt > 0 {
					arr = append(arr, -1)
					cnt--
				}
			}
		} else {
			if cnt > 0 {
				maxID++
				used_spaces = append(used_spaces, Space{maxID, len(arr), len(arr) + cnt - 1})
				for cnt > 0 {
					arr = append(arr, maxID)
					cnt--
				}
			}
		}
		freeSpace = !freeSpace
	}
	return arr, available_spaces, used_spaces, maxID
}
func checksum(arr []int) int {
	sum := 0
	for pos, id := range arr {
		if id == -1 {
			continue
		}
		sum += pos * id
	}
	return sum
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	defer util.Timer()()
	arr, _, _, _ := parseInput(strings.Split(input[0], ""))
	rearrage(arr)
	return strconv.Itoa(checksum(arr))
}
func rearrage(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left] != -1 {
			left++
			continue
		}
		if arr[right] == -1 {
			right--
			continue
		}
		arr[left] = arr[right]
		arr[right] = -1
		left++
		right--
	}
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	defer util.Timer()()
	arr, available_spaces, used_spaces, maxID := parseInput(strings.Split(input[0], ""))
	rearrage2(arr, available_spaces, used_spaces, maxID)

	arr2, _, _, _ := parseInput(strings.Split(input[0], ""))
	slow(arr2)
	fmt.Println(strconv.Itoa(checksum(arr)))
	fmt.Println(strconv.Itoa(checksum(arr2)))
	return strconv.Itoa(checksum(arr))
}

// segment tree??
func rearrage2(arr []int, available_spaces treeset.Set[Space], used_spaces []Space, maxID int) {
	// fmt.Printf("maxID=%v\n", maxID)
	// fmt.Println(available_spaces.String())
	// fmt.Println(used_spaces)
	// fmt.Println(arr)

	for i := len(used_spaces) - 1; i >= 0; i-- {
		currSpace := used_spaces[i]
		compact(arr, available_spaces, currSpace)
		// fmt.Printf("id=%v", currSpace.id)
		// fmt.Println(available_spaces.String())
		// fmt.Println(arr)
	}
}
func compact(arr []int, available_spaces treeset.Set[Space], currSpace Space) {
	index, freeSpace := available_spaces.Find(func(index int, c1 Space) bool {
		return (c1.available_space() >= currSpace.available_space())
	})
	if index < 0 {
		return
	}
	if freeSpace.start <= currSpace.start {
		available_spaces.Remove(freeSpace)
		for i := 0; i < currSpace.available_space(); i++ {
			arr[freeSpace.start+i] = currSpace.id
			arr[currSpace.start+i] = -1
		}
		//add back leftover
		//too fragmented?
		if currSpace.available_space() < freeSpace.available_space() {
			available_spaces.Add(Space{-1, freeSpace.end - (freeSpace.available_space() - currSpace.available_space()) + 1, freeSpace.end})
		}
	}
}

func slow(arr []int) {

	maxID := 0
	for _, id := range arr {
		if id > maxID {
			maxID = id
		}
	}

	for id := maxID; id > 0; id-- {
		fileStart, fileEnd := -1, -1

		for i := 0; i < len(arr); i++ {
			if arr[i] == id {
				if fileStart == -1 {
					fileStart = i
				}
				fileEnd = i
			}
		}

		if fileStart == -1 {
			continue
		}
		pageLen := fileEnd - fileStart + 1

		freeStart, freeLen := -1, 0
		for i := 0; i < fileStart; i++ {
			if arr[i] == -1 {
				if freeStart == -1 {
					freeStart = i
				}
				freeLen++
			} else {
				freeStart, freeLen = -1, 0
			}
			if freeLen == pageLen {
				break
			}
		}
		if freeStart != -1 && freeLen >= pageLen {
			for i := 0; i < pageLen; i++ {
				arr[freeStart+i] = id
			}

			for i := fileStart; i <= fileEnd; i++ {
				arr[i] = -1
			}
		}
	}
}

type Space struct {
	id    int
	start int
	end   int
}

func (m Space) available_space() int {
	return m.end - m.start + 1
}

// Custom comparator (sort by IDs)
func bySpace(c1, c2 Space) int {

	// Type assertion, program will panic if this is not respected
	// c1 := a.(Space)
	// c2 := b.(Space)

	if c1.available_space() == c2.available_space() {
		return c1.start - c2.start
	}
	return c1.available_space() - c2.available_space()
}
