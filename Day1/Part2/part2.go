package part2

import (
	shr "day1/shared"
)

func Solve() int {
	stLines := shr.ReadFileLines("part1/input.txt")
	intLines := shr.ConvertStringArrToIntArr(stLines)
	count := 0

	for ind := 1; ind < len(intLines)-2; ind++ {
		current := intLines[ind-1] + intLines[ind] + intLines[ind+1]
		next := intLines[ind] + intLines[ind+1] + intLines[ind+2]

		if next > current {
			count++
		}
	}

	return count
}
