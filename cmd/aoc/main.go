package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/valonsogit/aoc-2023/internal/advent"
	"github.com/valonsogit/aoc-2023/internal/days"
)

func main() {
	currDayArg := flag.Int("day", time.Now().Day(), "The day to execute, defaults to the current day of december, dont fall behind!")
	executeAllArg := flag.Bool("all", false, "Whether to execute all days instead of a single one.")
	testInputArg := flag.Bool("t", false, "Whether to use test inputs instead of full ones.")
	flag.Parse()

	days := []advent.Day{days.Day1{}, &days.Day2{}}
	if *executeAllArg {
		for i, v := range days {
			fmt.Printf("Day%v_P1: %v\n", i, v.Part1(advent.GetInput(i, *testInputArg)))
			fmt.Printf("Day%v_P2: %v\n", i, v.Part2(advent.GetInput(i, *testInputArg)))
		}
	} else {
		currentDayN := *currDayArg
		if len(days) < currentDayN {
			log.Fatalf("Missing day %v in the days array", currentDayN)
		}
		cD := days[currentDayN-1]

		p1 := cD.Part1(advent.GetInput(currentDayN, *testInputArg))

		if p1 != "" {
			fmt.Printf("Day%v_P1: %v\n", currentDayN, p1)
			p2 := cD.Part2(advent.GetInput(currentDayN, *testInputArg))
			if p2 != "" {
				fmt.Printf("Day%v_P2: %v\n", currentDayN, p2)
			}
		}
	}
}
