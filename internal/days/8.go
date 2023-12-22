package days

import (
	"fmt"
	"regexp"
	"strings"
)

type Day8 struct{}
type Node struct {
	L string
	R string
}

func (d Day8) Part1(input []string) string {
	nodeMap := map[string]Node{}

	for _, l := range input[2:] {
		label := strings.TrimSpace(strings.Split(l, "=")[0])
		lr := strings.Split(l, "=")[1]
		regex := regexp.MustCompile(`\((.{3}), (.{3})\)`)

		matches := regex.FindStringSubmatch(lr)
		l := matches[1]
		r := matches[2]

		nodeMap[label] = Node{
			L: l,
			R: r,
		}
	}
	// for k, v := range nodeMap {
	// 	// fmt.Printf("%v: %+v\n", k, v)
	// }
	currentPos := nodeMap["AAA"]
	steps := 0
	for {
		dir := rune(input[0][steps%len(input[0])])
		var next string
		if dir == 'L' {
			next = currentPos.L
		} else {
			next = currentPos.R
		}
		// fmt.Printf("Currently at %+v going %c -> %v\n", currentPos, dir, next)
		currentPos = nodeMap[next]
		steps++
		if next == "ZZZ" {
			break
		}
	}
	return fmt.Sprint(steps)
}
func FindDepth(nodeMap map[string]Node, directions, start string) int {
	steps := 0
	currentPos := nodeMap[start]
	for {
		dir := rune(directions[steps%len(directions)])
		var next string
		if dir == 'L' {
			next = currentPos.L
		} else {
			next = currentPos.R
		}
		// fmt.Printf("Currently at %+v going %c -> %v\n", currentPos, dir, next)
		currentPos = nodeMap[next]
		steps++
		if next[2] == 'Z' {
			return steps
		}
	}

}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func (d Day8) Part2G(input []string) string {
	nodeMap := map[string]Node{}

	for _, l := range input[2:] {
		label := strings.TrimSpace(strings.Split(l, "=")[0])
		lr := strings.Split(l, "=")[1]
		regex := regexp.MustCompile(`\((.{3}), (.{3})\)`)

		matches := regex.FindStringSubmatch(lr)
		l := matches[1]
		r := matches[2]

		nodeMap[label] = Node{
			L: l,
			R: r,
		}
	}
	startingNodes := make([]string, 0)
	for k := range nodeMap {
		// fmt.Printf("%v: %+v\n", k, v)
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}
	depths := make([]int, 0)
	for _, v := range startingNodes {
		fmt.Printf("Start %v has depth %v\n", v, FindDepth(nodeMap, input[0], v))
		depths = append(depths, FindDepth(nodeMap, input[0], v))
	}

	return fmt.Sprint(LCM(depths[0], depths[1], depths...))
}

func (d Day8) Part2(input []string) string {

	nodeMap := map[string]Node{}

	for _, l := range input[2:] {
		label := strings.TrimSpace(strings.Split(l, "=")[0])
		lr := strings.Split(l, "=")[1]
		regex := regexp.MustCompile(`\((.{3}), (.{3})\)`)

		matches := regex.FindStringSubmatch(lr)
		l := matches[1]
		r := matches[2]

		nodeMap[label] = Node{
			L: l,
			R: r,
		}
	}
	startingNodes := make([]Node, 0)
	for k, v := range nodeMap {
		// fmt.Printf("%v: %+v\n", k, v)
		if k[2] == 'A' {
			startingNodes = append(startingNodes, v)
		}
	}
	steps := 0
	for {
		dir := rune(input[0][steps%len(input[0])])
		finished := true
		for i := range startingNodes {

			currentPos := startingNodes[i]
			var next string
			if dir == 'L' {
				next = currentPos.L
			} else {
				next = currentPos.R
			}
			// fmt.Printf("(%v) Currently at %+v going %c -> %v\n",i, currentPos, dir, next)
			startingNodes[i] = nodeMap[next]
			if next[2] != 'Z' {
				finished = false
			}
		}
		steps++
		fmt.Printf("%v\r", steps)
		if finished {
			break
		}
	}

	return fmt.Sprint(steps)
}
