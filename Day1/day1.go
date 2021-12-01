package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	stLines := strings.Split(string(content), "\n")
	intLines := ConvertStringArrToIntArr(stLines)
	count := 0

	for ind := 1; ind < len(intLines); ind++ {
		if intLines[ind] > intLines[ind-1] {
			count += 1
		}
	}

	fmt.Println(count)
}

func ConvertStringArrToIntArr(stArr []string) []int {
	var intArr []int
	for _, stValue := range stArr {
		cleanStValue := strings.Replace(stValue, "\r", "", -1)
		intVar, err := strconv.Atoi(cleanStValue)

		if err != nil {
			fmt.Println(err)
		} else {
			intArr = append(intArr, intVar)
		}
	}
	return intArr
}
