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

	originString = os.Args[1]
	for _, char := range originString {
		_, err := strconv.Atoi(string(char))

		if err == nil {
			tempNumber += string(char)
		} else {
			if tempNumber != "" {
				number, _ := strconv.Atoi(tempNumber)
				numbers = append(numbers, number)
			}
			tempNumber = ""
		}
	}

	fmt.Println(len(numbers), "("+strings.Trim(fmt.Sprint(numbers), "[]")+")")
}
