package aoc

import (
	"fmt"
	"log"

	"github.com/valonsogit/aoc-2023/internal/util"
)

type Day interface {
	Part1([]string) string
	Part2([]string) string
}

var accessMap = make(map[int]int, 31)

func GetInput(day int, test bool) []string {
	if test {
		accessMap[day]++
		testPath := fmt.Sprintf("input/test/%v_%v.txt", day, accessMap[day])
		input, err := util.ReadFile(testPath)
		if err == nil {
			return input
		}
		fmt.Printf("Missing test input for day %v part %v, using full input\n", day, accessMap[day])
	}

	fullPath := fmt.Sprintf("input/full/%v.txt", day)
	input, err := util.ReadFile(fullPath)
	if err != nil {
		log.Panicf("Missing full input for day %v\n", day)
	}
	return input
}
