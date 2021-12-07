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

func ConvertStringToInt(str string) int {
	cleanStValue := strings.Replace(str, "\r", "", -1)
	intVar, err := strconv.Atoi(cleanStValue)

	if err != nil {
		fmt.Println(err)
	}

	return intVar
}
