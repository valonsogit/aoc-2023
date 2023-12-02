package main

import (
	"aoc/internal/aoc"
	"aoc/internal/days"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	currDayArg := flag.Int("day", time.Now().Day(), "The day to execute, defaults to the current day of december, dont fall behind!")
	executeAllArg := flag.Bool("all", false, "Whether to execute all days instead of a single one.")
	testInputArg := flag.Bool("t", false, "Whether to use test inputs instead of full ones.")
	flag.Parse()

	days := []aoc.Day{days.Day1{}, &days.Day2{}}
	if *executeAllArg {
		for i, v := range days {
			fmt.Printf("Day%v_P1: %v\n", i, v.Part1(aoc.GetInput(i, *testInputArg)))
			fmt.Printf("Day%v_P2: %v\n", i, v.Part2(aoc.GetInput(i, *testInputArg)))
		}
	} else {
		currentDayN := *currDayArg
		if len(days) < currentDayN {
			log.Fatalf("Missing day %v in the days array", currentDayN)
		}
		cD := days[currentDayN-1]

		p1 := cD.Part1(aoc.GetInput(currentDayN, *testInputArg))
		
		if p1 != "" {
			fmt.Printf("Day%v_P1: %v\n", currentDayN, p1)
			p2 := cD.Part2(aoc.GetInput(currentDayN, *testInputArg))
			if p2 != "" {
				fmt.Printf("Day%v_P2: %v\n", currentDayN, p2)
			}
		}
	}
}
