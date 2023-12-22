package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day9 struct{}

func Aux(d []int) []int {
	currArr := make([]int, len(d)-1)

	for i := len(d) - 1; i > 0; i-- {
		curr := d[i]
		prev := d[i-1]
		diff := curr - prev
		// fmt.Printf("%v - %v = %v\n", curr, prev, diff)
		currArr[i-1] = diff
	}
	return currArr
}
func Predict(data []int) int {
	multilevel := make([][]int, 0)
	multilevel = append(multilevel, data)
	currArr := data
	for {
		currArr = Aux(currArr)
		multilevel = append(multilevel, currArr)
		b := true
		for _, v := range currArr {
			if v != 0 {
				b = false
			}
		}
		if b {
			break
		}
	}
	fmt.Printf("%+v\n", multilevel)
	t := 0
	for i := range multilevel {
		fmt.Printf("%v += %v - %v\n", t, multilevel[len(multilevel)-1-i][0], t)
		t = multilevel[len(multilevel)-1-i][0] - t
	}
	return t
}
func (d Day9) Part1(input []string) string {
	t := 0
	for _, v := range input {
		nums := make([]int, 0)
		for _, n := range strings.Fields(v) {
			nP, _ := strconv.Atoi(n)
			nums = append(nums, nP)
		}
		x := Predict(nums)
		fmt.Println(x)
		t += x
	}
	return fmt.Sprint(t)
}

func (d Day9) Part2(input []string) string {
	return ":C"
}
