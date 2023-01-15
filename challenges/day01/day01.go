package day01

import (
	"aoc_go/util"
	"fmt"
)

func Part1(input string) {
	text, err := util.ReadFileAsText("test.txt")
	if err == nil {
		println(fmt.Sprintf("Err %v", err))
	} else {
		println(text)
	}
}
