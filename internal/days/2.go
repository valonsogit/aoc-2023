package days

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
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
func (d *Day2) Part1(input []string) string {
	games := make([]Game, 0)
	idReg := regexp.MustCompile(`Game (\d+)`)
	for lineN, v := range input {
		game := Game{}
		gameSubsetSplit := strings.Split(v, ":")
		idMatches := idReg.FindStringSubmatch(gameSubsetSplit[0])
		if len(idMatches) == 0 {
			log.Panicf("Uh oh, game at line %v doesnt have an ID\n", lineN)
		}
		parsedId, err := strconv.ParseUint(idMatches[1], 10, 32)
		if err != nil {
			log.Panicf("Uh oh, ID %v at line %v is not a valid uint\n", idMatches[1], lineN)

		}
		game.id = uint(parsedId)

		for _, round := range strings.Split(gameSubsetSplit[1], ";") {
			for _, play := range strings.Split(round, ",") {
				nColor := strings.Split(strings.TrimSpace(play), " ")
				switch nColor[1] {
				case "blue":
					{
						parsedN, err := strconv.ParseUint(nColor[0], 0, 32)
						if err != nil {
							log.Panicf("Uh oh, number of cubes %v at line %v is not a valid uint\n", nColor[0], lineN)
						}
						n := uint(parsedN)
						game.Blue += n
						if n > game.MaxBlue {
							game.MaxBlue = n
						}
					}
				case "red":
					{
						parsedN, err := strconv.ParseUint(nColor[0], 0, 32)
						if err != nil {
							log.Panicf("Uh oh, number of cubes %v at line %v is not a valid uint\n", nColor[0], lineN)
						}
						n := uint(parsedN)
						game.Red += n
						if n > game.MaxRed {
							game.MaxRed = n
						}
					}
				case "green":
					{
						parsedN, err := strconv.ParseUint(nColor[0], 0, 32)
						if err != nil {
							log.Panicf("Uh oh, number of cubes %v at line %v is not a valid uint\n", nColor[0], lineN)
						}
						n := uint(parsedN)
						game.Green += n
						if n > game.MaxGreen {
							game.MaxGreen = n
						}
					}
				default:
					{
						log.Panicf("Uh oh, encountered invalid color %v when playing game at line %v", nColor[1], lineN)

					}
				}
			}

		}

		games = append(games, game)
	}

	bag := Bag{Red: 12, Green: 13, Blue: 14}
	sumOfPossibleGames := uint(0)
	for _, g := range games {
		if g.MaxRed <= bag.Red && g.MaxBlue <= bag.Blue && g.MaxGreen <= bag.Green {
			sumOfPossibleGames += g.id
		}
	}
	fmt.Println(games)
	d.games = games
	return fmt.Sprintf("SumOfPossibleGames = %v\n", sumOfPossibleGames)
}
func (d *Day2) Part2(input []string) string {
	power := uint(0)
	for _, g := range d.games {
		power += (g.MaxBlue * g.MaxRed * g.MaxGreen)
	}
	return fmt.Sprintf("Power of all valid sets = %v", power)
}
