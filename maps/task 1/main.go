package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	var result = make([]int, 0)
	myMap := make(map[int]bool)

	for _, value := range arr {
		if _, found := myMap[value]; found {
			continue
		}
		result = append(result, value)
		myMap[value] = false
	}
	fmt.Println(result)
}
