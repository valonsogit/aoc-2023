package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Day5 struct{}

func (d Day5) Part1(input []string) string {
	transformationMap := make(map[int]int)
	seeds := strings.Fields(strings.Split(input[0], ":")[1])
	fmt.Printf("Starting seeds %v\n", seeds)
	transformedSeeds := make([]int, len(seeds))
	for i, s := range seeds {
		transformedSeeds[i], _ = strconv.Atoi(s)
	}
	origin := "seed"
	destination := "seed"
	for lineN, line := range input[1:] {
		if strings.Contains(line, ":") || lineN == len(input)-2 {
			// util.StructLog(transformationMap)
			fmt.Println()
			fmt.Println()
			fmt.Printf("From %v to %v\n", origin, destination)
			if lineN != len(input)-2 {
				origin = strings.Split(strings.Fields(line)[0], "-to-")[0]
				destination = strings.Split(strings.Fields(line)[0], "-to-")[1]
			}
			for i, seed := range transformedSeeds {
				if v, ok := transformationMap[seed]; ok {
					transformedSeeds[i] = v
					fmt.Printf("Transforming %v -> %v\n", seed, v)
				} else {
					fmt.Printf("Transforming %v -> %v\n", seed, seed)
				}
			}
			transformationMap = make(map[int]int)
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}

		numbers := strings.Fields(line)
		destIndex, _ := strconv.Atoi(numbers[0])
		sourceIndex, _ := strconv.Atoi(numbers[1])
		transformationRange, _ := strconv.Atoi(numbers[2])
		fmt.Printf("Transforming %v -> %v over %v\n", sourceIndex, destIndex, transformationRange)

		for _, v := range transformedSeeds {
			if v >= sourceIndex && v < sourceIndex+transformationRange {
				fmt.Printf("Transforming %v -> %v\n", v, destIndex+(v-sourceIndex))
				transformationMap[v] = destIndex + (v - sourceIndex)
			}
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Printf("From %v to %v\n", "seed", destination)
	for i, seed := range transformedSeeds {
		fmt.Printf("Transformed %v -> %v\n", seeds[i], seed)
	}
	lowest := math.MaxInt
	for _, v := range transformedSeeds {
		if v < lowest {
			lowest = v
		}
	}
	return fmt.Sprintf("The closest location is %v\n", lowest)

}

type Transform struct {
	source          string
	destination     string
	transformations []Transformation
}
type Transformation struct {
	destIndex   int
	sourceIndex int
	trRange     int
}

func findDistance(input []Transform, seed int) int {
	for _, transform := range input {
		for _, tr := range transform.transformations {
			if seed >= tr.sourceIndex && seed < tr.sourceIndex+tr.trRange {
				seed = tr.destIndex + (seed - tr.sourceIndex)
				break
			}
		}
	}
	return seed
}
func (d Day5) Part2(input []string) string {
	startTime := time.Now().Local().Unix()
	seedSection := strings.Fields(strings.Split(input[0], ":")[1])
	//Parse data
	transforms := make([]Transform, 0)
	var currentTransform *Transform
	for _, line := range input[1:] {
		if strings.Contains(line, ":") {
			if currentTransform != nil {

				transforms = append(transforms, *currentTransform)
			}
			currentTransform = &Transform{
				transformations: make([]Transformation, 0),
			}
			origin := strings.Split(strings.Fields(line)[0], "-to-")[0]
			destination := strings.Split(strings.Fields(line)[0], "-to-")[1]
			currentTransform.destination = destination
			currentTransform.source = origin

			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}

		numbers := strings.Fields(line)
		destIndex, _ := strconv.Atoi(numbers[0])
		sourceIndex, _ := strconv.Atoi(numbers[1])
		transformationRange, _ := strconv.Atoi(numbers[2])
		currentTransform.transformations = append(currentTransform.transformations, Transformation{
			destIndex:   destIndex,
			sourceIndex: sourceIndex,
			trRange:     transformationRange,
		})
	}
	transforms = append(transforms, *currentTransform)

	for _, t := range transforms {
		fmt.Printf("{\n\tsource: %v\n\tdestination: %v\n\ttransformations: [\n", t.source, t.destination)
		for _, tr := range t.transformations {
			fmt.Printf("\t\t{\n\t\t\tdestIndex: %v\n\t\t\tsourceIndex: %v\n\t\t\ttrRange: %v\n\t\t}\n", tr.destIndex, tr.sourceIndex, tr.trRange)
		}
		fmt.Printf("\t]\n}\n")
	}
	min := math.MaxInt
	minChannel := make(chan int, 30)
	processed := 0
	go func() {
		for v := range minChannel {
			if v < min {
				min = v
				fmt.Printf("\nMin is now %v\n", min)
			}
			processed += 1
			fmt.Printf("%v\r", processed)
		}
	}()
	
	var wg sync.WaitGroup
	total := 0
	for i := 0; i < len(seedSection)/2; i++ {
		start, _ := strconv.Atoi(seedSection[i*2])
		rang, _ := strconv.Atoi(seedSection[i*2+1])
		fmt.Printf("\nStarting section %v\n", i)
		if i != 0{
			endTime := time.Now().Local().Unix()
			fmt.Printf("Total time: %vH:%vM:%vS\n", (endTime-startTime)/3600, ((endTime-startTime)%3600)/60, ((endTime-startTime)%3600)%60)
		}
		for x := 0; x < rang; x++ {
			wg.Add(1)
			go func(x int) {
				defer wg.Done()
				total++
				minChannel <- findDistance(transforms, start+x)
			}(x)
			fmt.Printf("%v/%v (%f%%)\r", x, rang-1, float64(x)/float64(rang-1)*100)
		}
	}

	wg.Wait()
	close(minChannel)



	fmt.Println()
	return fmt.Sprintf("The closest location is %v\n", min)
}
