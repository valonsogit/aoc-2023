package days

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
	"slices"
	"sync"
	"time"

	colorize "github.com/gookit/color"
)

type Day10 struct{}
type Position struct {
	x, y int
}

type Pipe struct {
	Pos    Position
	Symbol rune
}

func (p Pipe) String() string {
	return fmt.Sprintf("%v %c", p.Pos, p.Symbol)
}

var used = make(map[Position]bool)

func PrintAreaAround(prev, pos Position, input []string) string {
	total := ""
	//Print first 30 lines or 30 previous lines of pos
	for y := pos.y - 30; y <= pos.y+30; y++ {
		for x := 0; x <= len(input[0]); x++ {
			if x == pos.x && y == pos.y {
				total += colorize.Green.Sprintf(string(input[y][x]))
			} else if x == prev.x && y == prev.y {
				total += colorize.Yellow.Sprintf(string(input[y][x]))
				used[Position{x, y}] = true
			} else if used[Position{x, y}] {
				total += colorize.Red.Sprintf(string(input[y][x]))
			} else if x >= 0 && x < len(input[0]) && y >= 0 && y < len(input) {
				total += string(input[y][x])
			}
		}
		if y >= 0 && y < len(input) {
			total += "\n"
		}
	}
	// fmt.Println()

	return total

}

var frames = make([]*image.Paletted, 0)
var discarded = make([]Position, 0)
var usedQueue = make([]Position, 0)

var maxQueue = 255

func CreateFrame(pos Position, input []string) {

	img := image.NewPaletted(image.Rect(0, 0, len(input[0])*4, len(input)*4), palette.Plan9)
	//Iterate all input
	var wg sync.WaitGroup
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				if slices.ContainsFunc(usedQueue, func(p Position) bool { return p.x == x && p.y == y }) {
					it := slices.IndexFunc(usedQueue, func(p Position) bool { return p.x == x && p.y == y })
					factor := 255 / maxQueue
					for i := 0; i < 4; i++ {
						for j := 0; j < 4; j++ {

							img.Set(x*4+i, y*4+j, color.RGBA{
								R: 255 - uint8(maxQueue-len(usedQueue)+it)*uint8(factor),
								G: uint8(maxQueue-len(usedQueue)+it) * uint8(factor),
								B: 0,
								A: 255,
							})
						}
					}
					// fmt.Printf("Set %v to %v\n", Position{x, y}, color.RGBA{0, uint8(255 - it), uint8(it), 255})

				} else if slices.ContainsFunc(discarded, func(p Position) bool { return p.x == x && p.y == y }) {
					for i := 0; i < 4; i++ {
						for j := 0; j < 4; j++ {
							img.Set(x*4+i, y*4+j, color.RGBA{255, 0, 0, 255})
						}
					}

				}
			}(x, y)
		}
	}

	wg.Wait()
	frames = append(frames, img)
}
func NextPosition(prev, curr Pipe) Position {
	// | is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.

	diff := Position{curr.Pos.x - prev.Pos.x, curr.Pos.y - prev.Pos.y}
	// fmt.Println(diff)
	if diff.y == 0 {
		switch curr.Symbol {
		case '-':
			{
				if curr.Pos.x > prev.Pos.x {
					return Position{curr.Pos.x + 1, curr.Pos.y}
				} else {
					return Position{curr.Pos.x - 1, curr.Pos.y}
				}
			}
		case 'L':
			{
				return Position{curr.Pos.x, curr.Pos.y - 1}
			}
		case 'J':
			{
				return Position{curr.Pos.x, curr.Pos.y - 1}
			}
		case '7':
			{
				return Position{curr.Pos.x, curr.Pos.y + 1}
			}
		case 'F':
			{
				return Position{curr.Pos.x, curr.Pos.y + 1}
			}
		}
	} else if diff.x == 0 {
		switch curr.Symbol {
		case '|':
			{
				if curr.Pos.y > prev.Pos.y {
					return Position{curr.Pos.x, curr.Pos.y + 1}
				} else {
					return Position{curr.Pos.x, curr.Pos.y - 1}
				}
			}
		case 'L':
			{
				return Position{curr.Pos.x + 1, curr.Pos.y}
			}
		case 'J':
			{
				return Position{curr.Pos.x - 1, curr.Pos.y}
			}
		case '7':
			{
				return Position{curr.Pos.x - 1, curr.Pos.y}
			}
		case 'F':
			{
				return Position{curr.Pos.x + 1, curr.Pos.y}
			}
		}
	}
	return Position{-1, -1}

}

func (d Day10) Part1(input []string) string {

	var pos Position
out:
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				pos = Position{x, y}
				break out
			}
		}

	}
	prev := Pipe{pos, 'S'}
	manualNext := Position{pos.x, pos.y + 1}
	curr := Pipe{manualNext, rune(input[manualNext.y][manualNext.x])}
	iter := 1
	for {
		usedQueue = append(usedQueue, curr.Pos)
		if len(usedQueue) > maxQueue {
			discarded = append(discarded, usedQueue[0])
			usedQueue = usedQueue[1:]
		}
		if iter%5 == 0 {
			CreateFrame(curr.Pos, input)
		}
		fmt.Println(iter)
		fmt.Print("\033[H\033[2J", PrintAreaAround(prev.Pos, curr.Pos, input))
		// fmt.Println(prev, curr)
		next := NextPosition(prev, curr)
		iter++
		if next.x == -1 || next.y == -1 {
			break
		}
		prev = curr
		curr = Pipe{next, rune(input[next.y][next.x])}
		time.Sleep(1 * time.Millisecond)

	}
	CreateFrame(curr.Pos, input)

	fmt.Printf("Finished at %v after %v iterations\n", curr, iter)

	f, err := os.OpenFile("rgb5.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	delays := make([]int, len(frames))
	delays[len(frames)-1] = 200
	gif.EncodeAll(f, &gif.GIF{
		Image: frames,
		Delay: delays,
	})
	// c := 0
	// last := frames[len(frames)-1]
	//

	return fmt.Sprint(iter / 2)
}

func (d Day10) Part2(input []string) string {
	return ""
}
