package part1

import (
	shr "day1/shared"
)

func Solve() int {
	stLines := shr.ReadFileLines("part1/input.txt")
	intLines := shr.ConvertStringArrToIntArr(stLines)
	count := 0

	for ind := 1; ind < len(intLines); ind++ {
		if intLines[ind] > intLines[ind-1] {
			count += 1
		}
	}

	return count
}
