package day02

import (
	"advent/utils"
	"os"
	"strconv"
	"strings"
)

// levelIsSafe checks if numbers differ by 1-3 in a consistently ascending or descending sequence
func levelIsSafe(level []int) bool {
	if len(level) < 2 {
		return true
	}

	ascending := level[0] < level[1]
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]

		if ascending && (diff < 1 || diff > 3) {
			return false
		}

		if !ascending && (diff > -1 || diff < -3) {
			return false
		}
	}
	return true
}

func Solve() int {
	input, err := os.ReadFile("./days/day02/input.txt")
	utils.Check(err)

	reports := strings.Split(string(input), "\n")
	validReports := 0

	for _, report := range reports {

		split := strings.Split(report, " ")
		levels := make([]int, 0)

		for _, level := range split {
			parsed, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}

			levels = append(levels, int(parsed))
		}

		if !levelIsSafe(levels) {
			continue
		}

		validReports++
	}

	return validReports

}
