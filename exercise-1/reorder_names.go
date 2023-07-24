package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var output string
	args := os.Args[1:]
	lastElement := args[len(args)-1]

	switch lastElement {
	case "US":
		output = strings.Join(args[0:len(args)-1], " ")
	case "VN":
		output = args[1] + " " + args[0]
	default:
		output = "Country code invalid"
	}

	fmt.Println(output)
}
