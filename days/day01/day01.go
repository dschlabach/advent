package day01

import (
	"advent/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() int {
	input, err := os.ReadFile("./days/day01/input.txt")
	utils.Check(err)

	lines := strings.Split(string(input), "\n")

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, line := range lines {
		// Get the values of both numbers in the line
		nums := strings.Split(line, "   ")

		a, _ := strconv.ParseInt(nums[0], 10, 0)
		b, _ := strconv.ParseInt(nums[1], 10, 0)

		list1 = append(list1, int(a))
		list2 = append(list2, int(b))
	}

	// Sort each array ascending
	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	distance := 0

	// Go through each pair and add them up
	for idx, val := range list1 {
		diff := val - list2[idx]

		if diff <= 0 {
			diff = -diff
		}

		distance += int(diff)
	}

	return distance
}
