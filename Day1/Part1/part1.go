package part1

import (
	shr "day1/shared"
	"fmt"
	"io/ioutil"
	"strings"
)

func Solve() int {

	content, err := ioutil.ReadFile("part1/input.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	stLines := strings.Split(string(content), "\n")
	intLines := shr.ConvertStringArrToIntArr(stLines)
	count := 0

	for ind := 1; ind < len(intLines); ind++ {
		if intLines[ind] > intLines[ind-1] {
			count += 1
		}
	}

	return count
}
