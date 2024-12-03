package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ----- Hand -----
type Hand struct {
	rank  int
	cards []string
	bet   int
}

func (hand *Hand) CalculateNumberOfWins() int {
	wins := 0

	return wins
}

// -----------------------

func main() {
	fmt.Println("AoC 2023")
	aoc1207()
}

func aoc1207() {
	lines := getLines("./input.txt")
	processInput(lines)

	//var result = calculateWinningCombinations(processInput(lines))
	//var result2 = calculateWinningCombinations(processInput2(lines))
	//fmt.Println("Day 06 Part 1 Result: ", result)
	//fmt.Println("Day 06 Part 2 Result: ", result2)
}

//func orderHands(hands []Hand) []Hand {
//
//}

func processInput(hands []string) []Hand {

	rounds := make([]Hand, len(hands))

	for i := 0; i < len(hands); i++ {
		hand := strings.Fields(strings.Split(hands[i], " ")[0])
		bet, _ := strconv.Atoi(strings.Split(hands[i], " ")[1])

		rounds[i] = Hand{0, hand, bet}
	}

	return rounds
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}
