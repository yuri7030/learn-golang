package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var originString string
	var tempNumber string
	var numbers []int
	numberMap := make(map[int]bool)

	originString = os.Args[1]
	for _, char := range originString {
		_, err := strconv.Atoi(string(char))

		if err == nil {
			tempNumber += string(char)
		} else {
			if tempNumber != "" {
				number, _ := strconv.Atoi(tempNumber)
				if numberMap[number] == false {
					numberMap[number] = true
				}
			}
			tempNumber = ""
		}
	}

	for key, _ := range numberMap {
		numbers = append(numbers, key)
	}

	fmt.Println(len(numbers), "("+strings.Trim(fmt.Sprint(numbers), "[]")+")")
}
