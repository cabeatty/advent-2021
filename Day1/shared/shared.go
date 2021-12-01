package shared

import (
	"fmt"
	"strconv"
	"strings"
)

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
