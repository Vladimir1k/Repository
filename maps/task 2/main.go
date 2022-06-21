package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var (
		result string
		input  = "1 9 3 4 -5"
	)
	inputArr := strings.Split(input, " ")
	min2, _ := strconv.Atoi(inputArr[0])
	min := int32(min2)
	max := min
	if len(inputArr) == 1 {
		result = inputArr[0]
	} else {

		for _, value := range inputArr {
			count, _ := strconv.Atoi(value)
			switch {
			case int32(count) < min:
				min = int32(count)
			case int32(count) > max:
				max = int32(count)
			}

		}
		result = fmt.Sprintf("%d %d", max, min)
	}

	fmt.Println(result)
}

///
