package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ----- Race -----
type Race struct {
	id             int
	totalTime      int
	recordDistance int
}

func (race *Race) CalculateNumberOfWins() int {
	wins := 0

	for t := 0; t < race.totalTime; t++ {
		if t != 0 {
			raceTime := race.totalTime - t
			totalDistance := raceTime * t

			if totalDistance > race.recordDistance {
				wins++
			}
		}
	}

	return wins
}

// -----------------------

func main() {
	fmt.Println("AoC 2023")
	aoc1206()
}

func aoc1206() {
	lines := getLines("./input.txt")

	var result = calculateWinningCombinations(processInput(lines))
	var result2 = calculateWinningCombinations(processInput2(lines))
	fmt.Println("Day 06 Part 1 Result: ", result)
	fmt.Println("Day 06 Part 2 Result: ", result2)
}

func calculateWinningCombinations(races []Race) int {
	totalWaysToWin := 1

	for _, race := range races {
		totalWaysToWin *= race.CalculateNumberOfWins()
	}

	return totalWaysToWin
}

func processInput(races []string) []Race {

	times := strings.Split(strings.TrimSpace(strings.Replace(races[0][12:], "  ", "", -1)), " ")
	distances := strings.Split(strings.TrimSpace(strings.Replace(races[1][11:], "  ", "", -1)), " ")

	boatRaces := make([]Race, len(times))

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(strings.TrimSpace(times[i]))
		distance, _ := strconv.Atoi(strings.TrimSpace(distances[i]))

		boatRaces[i] = Race{i + 1, time, distance}
	}

	return boatRaces
}

func processInput2(races []string) []Race {

	times := strings.Split(strings.TrimSpace(strings.Replace(races[0][12:], "  ", "", -1)), " ")
	distances := strings.Split(strings.TrimSpace(strings.Replace(races[1][11:], "  ", "", -1)), " ")

	time, _ := strconv.Atoi(strings.Join(times, ""))
	distance, _ := strconv.Atoi(strings.Join(distances, ""))

	fmt.Println("Time:", time)
	fmt.Println("Distance", distance)

	race := []Race{{1, time, distance}}

	return race
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}
