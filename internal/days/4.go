package days

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

type Day4 struct{}

func (d Day4) Part1(input []string) string {
	t := 0
	for _, line := range input {
		data := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := strings.Fields(data[0])
		ownedNumbers := strings.Fields(data[1])
		won := -1
		for _, v := range ownedNumbers {
			if slices.Contains(winningNumbers, v) {
				// fmt.Printf("%v is a winning number\n", v)
				won += 1
			}
		}
		if won != -1 {
			// fmt.Printf("Card %v is worth %v points\n", lineN, math.Pow(2, float64(won)))
			t += int(math.Pow(2, float64(won)))
		}
	}
	return fmt.Sprint(t)
}

func (d Day4) Part2(input []string) string {
	totalCards := make(map[int]int)
	for lineN, line := range input {
		for i := 0; i < totalCards[lineN]+1; i += 1 {
			data := strings.Split(strings.Split(line, ":")[1], "|")
			winningNumbers := strings.Fields(data[0])
			ownedNumbers := strings.Fields(data[1])
			won := 0
			for _, v := range ownedNumbers {
				if slices.Contains(winningNumbers, v) {
					// fmt.Printf("%v is a winning number\n", v)
					won += 1
				}
			}
			for i := 1; i < won+1; i += 1 {
				totalCards[lineN+i] += 1
			}
		}
		totalCards[lineN] = totalCards[lineN] + 1

	}
	t := 0
	for i, v := range totalCards {
		fmt.Printf("Card %v has %v copies\n", i+1, v)
		t += v
	}
	return fmt.Sprint(t)
}
