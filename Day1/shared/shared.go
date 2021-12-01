package shared

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileLines(filePath string) []string {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err.Error())
	}

	return strings.Split(string(content), "\n")
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
