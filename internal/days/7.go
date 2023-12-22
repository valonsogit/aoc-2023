package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day7 struct{}
type Hand struct {
	play  string
	bid   int
	score int
	t     int
}

func (d Day7) Part1(input []string) string {
	scoring := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	hands := make([]*Hand, 0)

	for _, line := range input {
		b, _ := strconv.Atoi(strings.Split(line, " ")[1])
		hands = append(hands, &Hand{
			play: strings.Split(line, " ")[0],
			bid:  b,
		})
	}
	for _, h := range hands {
		cCount := map[rune]int{}
		for i, c := range h.play {
			h.score += scoring[c] * int(math.Pow(13, float64(4-i)))
			cCount[c]++
		}

		vals := make([]int, 0, len(cCount))
		for _, v := range cCount {
			vals = append(vals, v)
		}
		sort.Ints(vals)
		sort.Sort(sort.Reverse(sort.IntSlice(vals)))

		s := 0
		for i, v := range vals {
			s += v * int(math.Pow(6, float64(5-i)))
		}

		h.t = s

	}

	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]
		if h1.t < h2.t {
			return true
		} else if h1.t > h2.t {
			return false
		} else {
			return h1.score < h2.score
		}
	})
	totalScore := 0
	for i, h := range hands {
		totalScore += (i + 1) * h.bid
	}
	return fmt.Sprint(totalScore)
}

func (d Day7) Part2(input []string) string {
	scoring := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}
	hands := make([]*Hand, 0)

	for _, line := range input {
		b, _ := strconv.Atoi(strings.Split(line, " ")[1])
		hands = append(hands, &Hand{
			play: strings.Split(line, " ")[0],
			bid:  b,
		})
	}
	for _, h := range hands {
		cCount := map[rune]int{}
		j := 0
		for i, c := range h.play {
			// fmt.Printf("%c at idx %v score: %v * %v = %v\n", c, 5-i, scoring[c], int(math.Pow(12, float64(4-i))), scoring[c]*int(math.Pow(12, float64(4-i))))
			if c == 'J' {
				j++
			} else {
				h.score += scoring[c] * int(math.Pow(13, float64(4-i)))
				cCount[c]++
			}
		}

		vals := make([]int, 5)
		for _, v := range cCount {
			vals = append(vals, v)
		}

		sort.Ints(vals)
		sort.Sort(sort.Reverse(sort.IntSlice(vals)))
		vals[0] += j
		s := 0
		for i, v := range vals {
			s += v * int(math.Pow(6, float64(5-i)))
		}

		h.t = s

		// h.score += int(math.Pow(10, 7-float64(len(cCount))))
	}

	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]
		if h1.t < h2.t {
			return true
		} else if h1.t > h2.t {
			return false
		} else {
			return h1.score < h2.score
		}
	})
	totalScore := 0
	for i, h := range hands {
		// fmt.Printf("Hand %v with bid %v type %v: has a score %v,  %v * %v =  %v\n", h.play, h.bid, h.t, h.score, (i + 1), h.bid, (i+1)*h.bid)
		totalScore += (i + 1) * h.bid
	}
	return fmt.Sprint(totalScore)
}
