package main

import (
	"aoc_go/util"
	"fmt"
)

func main() {
	input, err := util.ReadFileAsText("input/test.txt")
	if err != nil {
		isEmpty := len(input) == 0
		fmt.Println("Input: ", isEmpty)
		fmt.Println(err)
	}
	if len(input) > 0 {
		println(input)
	}
}
