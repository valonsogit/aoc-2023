package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Day2 struct {
	games []Game
}
type Bag struct {
	Red   uint
	Blue  uint
	Green uint
}
type Game struct {
	id       uint
	Red      uint
	Blue     uint
	Green    uint
	MaxRed   uint
	MaxBlue  uint
	MaxGreen uint
}

func (g Game) String() string {
	return fmt.Sprintf("{\n\tid: %v\n\tRed: %v\n\tBlue: %v\n\tGreen: %v\n\tMaxRed: %v\n\tMaxBlue: %v\n\tMaxGreen: %v\n}", g.id, g.Red, g.Blue, g.Green, g.MaxRed, g.MaxBlue, g.MaxGreen)
}

func ParseLine(v string) Game {
	game := Game{}
	gameSubsetSplit := strings.Split(v, ":")
	idReg := regexp.MustCompile(`Game (\d+)`)

	idMatches := idReg.FindStringSubmatch(gameSubsetSplit[0])

	parsedId, _ := strconv.ParseUint(idMatches[1], 10, 32)

	game.id = uint(parsedId)

	for _, round := range strings.Split(gameSubsetSplit[1], ";") {
		for _, play := range strings.Split(round, ",") {

			nColor := strings.Split(strings.TrimSpace(play), " ")

			switch nColor[1] {
			case "blue":
				{
					parsedN, _ := strconv.ParseUint(nColor[0], 0, 32)

					n := uint(parsedN)
					game.Blue += n
					if n > game.MaxBlue {
						game.MaxBlue = n
					}
				}
			case "red":
				{
					parsedN, _ := strconv.ParseUint(nColor[0], 0, 32)

					n := uint(parsedN)
					game.Red += n
					if n > game.MaxRed {
						game.MaxRed = n
					}
				}
			case "green":
				{
					parsedN, _ := strconv.ParseUint(nColor[0], 0, 32)

					n := uint(parsedN)
					game.Green += n
					if n > game.MaxGreen {
						game.MaxGreen = n
					}
				}
			default:
				{

				}
			}
		}

	}
	return game
}

func (d Day2) Part1(input []string) string {
	gamesChannel := make(chan Game, len(input))

	var wg sync.WaitGroup

	for _, v := range input {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			gamesChannel <- ParseLine(v)
		}(v)
	}
	go func() {
		wg.Wait()
		close(gamesChannel)
	}()

	bag := Bag{Red: 12, Green: 13, Blue: 14}

	sumOfPossibleGames := uint(0)

	for g := range gamesChannel {
		if g.MaxRed <= bag.Red && g.MaxBlue <= bag.Blue && g.MaxGreen <= bag.Green {
			sumOfPossibleGames += g.id
		}
	}

	return fmt.Sprint(sumOfPossibleGames)
}
func (d Day2) Part2(input []string) string {
	power := uint(0)
	// for _, g := range d.games {
	// 	power += (g.MaxBlue * g.MaxRed * g.MaxGreen)
	// }
	return fmt.Sprintf("Power of all valid sets = %v", power)
}
