package day02

import (
	"advent/utils"
	"os"
	"strconv"
	"strings"
)

// levelIsSafe checks if numbers differ by 1-3 in a consistently ascending or descending sequence
func levelIsSafe(level []int) bool {
	ascending := level[0] < level[1]

	for i := 0; i < len(level)-1; i++ {
		if ascending {
			if level[i+1] < level[i]+1 || level[i+1] > level[i]+3 {
				return false
			}
		}

		if !ascending {
			if level[i+1] > level[i]-1 || level[i+1] < level[i]-3 {
				return false
			}
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
			parsed, err := strconv.ParseInt(level, 10, 0)
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
