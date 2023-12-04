package main

import (
	"fmt"

	"github.com/valonsogit/aoc-2023/internal/days"
	"github.com/valonsogit/aoc-2023/internal/util"
)

func main() {
	in, _ := util.ReadFile("input/full/2.txt")
	fmt.Println(days.Day2{}.Part1(in))
}
