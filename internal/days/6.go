package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day6 struct{}
type Race struct {
	Time     int
	Distance int
}

func minMax(totalTime, Y float64) (minR, maxR float64) {
	rightSide := math.Sqrt(math.Pow(totalTime, 2) - 4*Y)
	min := (totalTime + rightSide) / 2
	max := (totalTime - rightSide) / 2
	return min, max
}
func (d Day6) Part1(input []string) string {
	times := strings.Fields(strings.TrimSpace(strings.Split(input[0], ":")[1]))
	distances := strings.Fields(strings.TrimSpace(strings.Split(input[1], ":")[1]))
	races := make([]Race, 0)
	for i := range times {
		d, _ := strconv.Atoi(distances[i])
		t, _ := strconv.Atoi(times[i])

		races = append(races, Race{
			Time:     t,
			Distance: d,
		})
	}
	acc := 1
	for _, r := range races {
		totalTime := r.Time
		Y := r.Distance
		min, max := minMax(float64(totalTime), float64(Y))
		fmt.Printf("min: %v | max: %v -> %v\n", min, max, int(math.Ceil(min)-math.Ceil(max)))

		acc *= int(math.Ceil(min) - math.Ceil(max+0.000001))
	}
	return fmt.Sprint(acc)
}

func (d Day6) Part2(input []string) string {
	dist, _ := strconv.Atoi(strings.Join(strings.Fields(strings.TrimSpace(strings.Split(input[1], ":")[1])), ""))
	t, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(input[0], ":")[1]), ""))
	fmt.Printf("t:%v d:%v", t, dist)
	totalTime := t
	Y := dist

	min, max := minMax(float64(totalTime), float64(Y))
	fmt.Printf("min: %v | max: %v -> %v\n", min, max, int(math.Ceil(min)-math.Ceil(max)))

	acc := int(math.Ceil(min) - math.Ceil(max+0.000001))
	return fmt.Sprint(acc)
}
