package days

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"github.com/gookit/color"
	"github.com/valonsogit/aoc-2023/internal/util"
)

type Gear struct {
	Ratio          int
	ConnectedParts int
}
type Day3 struct {
	gearMap map[int]map[int]*Gear
}

func (d *Day3) LookAround(input []string, lineN, start, end int) bool {
	checkedLines := []int{}
	if lineN > 0 {
		checkedLines = append(checkedLines, lineN-1)
		// fmt.Printf("Checking lines %v %v %v\n", checkedLines[0], checkedLines[1], checkedLines[2])
	}
	checkedLines = append(checkedLines, lineN)

	if lineN < len(input)-1 {
		checkedLines = append(checkedLines, lineN+1)

	}
	red := color.FgRed.Render
	green := color.FgGreen.Render

	strTotal := ""
	help := false
	for _, l := range checkedLines {

		for i := max(start-1, 0); i < min(end+1, len(input[lineN])); i++ {
			char := rune(input[l][i])
			if char != '.' && !unicode.IsDigit(char) {
				strTotal += green(fmt.Sprintf("%c", char))
				help = true
				if char == '*' {
					parsedN, _ := strconv.ParseInt(input[lineN][start:end], 10, 32)
					if d.gearMap[i] == nil {
						d.gearMap[i] = make(map[int]*Gear)
					}
					if _, ok := d.gearMap[i][l]; !ok {
						d.gearMap[i][l] = &Gear{
							Ratio:          1,
							ConnectedParts: 0,
						}
					}
					d.gearMap[i][l].Ratio = d.gearMap[i][l].Ratio * int(parsedN)
					d.gearMap[i][l].ConnectedParts += 1

				}
			} else {
				strTotal += red(fmt.Sprintf("%c", char))
			}
		}
		strTotal += "\n"
	}
	if !help {
		fmt.Print(strTotal)
	}
	return help

}

func (d *Day3) Part1(input []string) string {
	d.gearMap = make(map[int]map[int]*Gear)

	numberRegex := regexp.MustCompile(`(\d+)`)
	total := 0
	for lineN, line := range input {
		numberIndexes := numberRegex.FindAllStringIndex(line, -1)
		for _, v := range numberIndexes {
			// fmt.Printf("N: %v index %v/%v (%c/%c)\n", line[v[0]:v[1]], v[0], v[1], line[v[0]], line[v[1]-1])
			if d.LookAround(input, lineN, v[0], v[1]) {
				// fmt.Printf("N: %v has adjacent symbol\n", line[v[0]:v[1]])
				parsedInt, _ := strconv.ParseInt(line[v[0]:v[1]], 10, 32)
				total += int(parsedInt)
			} else {
				println()
				println()
			}

		}
	}
	return fmt.Sprint(total)
}

func (d *Day3) Part2(input []string) string {
	totalRatio := 0
	for i := range d.gearMap {
		for j := range d.gearMap[i] {
			gear := *d.gearMap[i][j]
			util.StructLog(gear)
			if gear.ConnectedParts == 2 {
				totalRatio += gear.Ratio
			}
		}
	}
	return fmt.Sprint(totalRatio)
}
