package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC 2023")
	aoc1202()
}

func aoc1202() {
	lines := getLines("./input.txt")

	var result = gameSum(lines)
	var result2 = fewestCubes(lines)
	fmt.Println("Day 01 Part 1 Result: ", result)
	fmt.Println("Day 01 Part 2 Result: ", result2)
}

func gameSum(lines []string) int {
	var gameSumVal = 0
	var red, green, blue = 12, 13, 14

	for _, line := range lines {
		validGame := true
		game := strings.Split(line, ":")
		id, _ := strconv.Atoi((strings.Split(game[0], " "))[1])
		rounds := strings.Split(game[1], ";")

		for i := 0; i < len(rounds); i++ {
			cubeVals := strings.Split(rounds[i], ",")

			for x := 0; x < len(cubeVals); x++ {
				cubes := strings.Split(cubeVals[x], " ")
				color := cubes[2]
				numColor, _ := strconv.Atoi(cubes[1])

				if color == "red" {
					if numColor > red {
						validGame = false
						break
					}

				} else if color == "blue" {
					if numColor > blue {
						validGame = false
						break
					}

				} else if color == "green" {
					if numColor > green {
						validGame = false
						break
					}
				}
			}
		}

		if validGame {
			gameSumVal = gameSumVal + id
		} else {
			fmt.Println("Game", id, "is invalid (", line, ")")
		}

	}

	return gameSumVal
}

func fewestCubes(lines []string) int {
	var totalGamePower = 0

	for _, line := range lines {
		var red, green, blue = 0, 0, 0

		game := strings.Split(line, ":")
		id, _ := strconv.Atoi((strings.Split(game[0], " "))[1])
		rounds := strings.Split(game[1], ";")

		fmt.Println("Game: ", game)
		fmt.Println("    ID: ", id)
		fmt.Println("    Rounds: ", rounds)

		for i := 0; i < len(rounds); i++ {
			cubeVals := strings.Split(rounds[i], ",")

			fmt.Println("        Round ", i+1, ": ", cubeVals)

			for x := 0; x < len(cubeVals); x++ {
				cubes := strings.Split(cubeVals[x], " ")
				color := cubes[2]
				numColor, _ := strconv.Atoi(cubes[1])

				fmt.Println("            ", color, " = ", numColor)

				if color == "red" {
					if numColor > red {
						red = numColor
					}

				} else if color == "blue" {
					if numColor > blue {
						blue = numColor
					}

				} else if color == "green" {
					if numColor > green {
						green = numColor
					}
				}
			}
		}

		fmt.Println("                For game ", id, " the fewest number of cubes of each color are red = ", red, ", green = ", green, ", blue = ", blue, ".")

		gamePower := red * green * blue
		totalGamePower = totalGamePower + gamePower

	}

	return totalGamePower
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}
