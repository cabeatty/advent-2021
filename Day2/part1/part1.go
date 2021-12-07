package part1

import (
	shr "day2/shared"
	"strings"
)

func Solve() int {
	var input = shr.ReadFileLines("part1/input.txt")
	var horizontal = 0
	var vertical = 0

	for _, movement := range input {

		var directions = strings.Split(string(movement), " ")

		var distance = shr.ConvertStringToInt(directions[1])
		var direction = directions[0]

		if direction == "forward" {
			horizontal = horizontal + distance
		} else if direction == "up" {
			vertical = vertical - distance
		} else if direction == "down" {
			vertical = vertical + distance
		}
	}

	return horizontal * vertical
}
