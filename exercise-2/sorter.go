package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var arrayInt []int
	var arrayString []string

	dataType := os.Args[1]
	array := os.Args[2:]

	for _, i := range array {
		j, err := strconv.Atoi(i)
		if err != nil {
			arrayString = append(arrayString, i)
		} else {
			arrayInt = append(arrayInt, j)
		}
	}

	switch dataType {
	case "-int":
		sort.Ints(arrayInt)
		printArr(arrayInt)
	case "-string":
		sort.Strings(arrayString)
		printArr(arrayString)
	default:
		sort.Ints(arrayInt)
		sort.Strings(arrayString)
	}

	// if len(arrayInt) > 0 {
	// 	for _, i := range arrayInt {
	// 		output += fmt.Sprintf("%d ", i)
	// 	}
	// }

	// if len(arrayString) > 0 {
	// 	output += strings.Join(arrayString, " ")
	// }

	// printArr(output)
}

func printArr(arr ...any) {
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}
