package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	var diagnostics = ReadFileLines("input.txt")
	var gamma = ""
	var epsilon = ""

	for i := 0; i < len(diagnostics[0])-1; i++ {
		var zeroCount = 0
		var oneCount = 0

		for _, diagnostic := range diagnostics {
			if string(diagnostic[i]) == "0" {
				zeroCount++
			} else {
				oneCount++
			}
		}

		if zeroCount > oneCount {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaVal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonVal, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaVal * epsilonVal)
}

func part2() int {
	var diagnostics = ReadFileLines("input.txt")
	var o2rating = part2_o2(diagnostics)
	var co2rating = part2_co2(diagnostics)

	o2int, _ := strconv.ParseInt(o2rating, 2, 64)
	co2int, _ := strconv.ParseInt(co2rating, 2, 64)
	rating := o2int * co2int

	return int(rating)
}

func part2_o2(diagnostics []string) string {
	bitSize := len(diagnostics[0]) - 1

	for i := 0; i <= bitSize; i++ {
		var ones []string
		var zeros []string

		for _, diagnostic := range diagnostics {
			thisBit := diagnostic[i : i+1]

			if thisBit == "0" {
				zeros = append(zeros, diagnostic)
			} else {
				ones = append(ones, diagnostic)
			}
		}

		if len(ones) < len(zeros) {
			diagnostics = zeros
		} else {
			diagnostics = ones
		}

		if len(diagnostics) == 1 {
			return diagnostics[0]
		}
	}

	return "nan"
}

func part2_co2(diagnostics []string) string {
	bitSize := len(diagnostics[0]) - 1

	for i := 0; i < bitSize; i++ {
		var ones []string
		var zeros []string

		for _, diagnostic := range diagnostics {
			thisBit := diagnostic[i : i+1]

			if thisBit == "0" {
				zeros = append(zeros, diagnostic)
			} else {
				ones = append(ones, diagnostic)
			}
		}

		if len(ones) < len(zeros) {
			diagnostics = ones
		} else {
			diagnostics = zeros
		}

		if len(diagnostics) == 1 {
			return diagnostics[0]
		}
	}

	return "nan"
}

func ReadFileLines(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	return txtlines
}
