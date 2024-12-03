package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var symbols = map[string]string{"-": "", "+": "", "@": "", "#": "", "$": "", "%": "", "&": "", "*": "", "=": "", "/": ""}

type Point struct {
	x, y int
}

func main() {
	fmt.Println("AoC 2023")
	aoc1202()
}

func aoc1202() {
	lines := getLines("./input.txt")

	var partNumSum, gearRatioSum = sumPartNumbers(lines)
	//var result2 = fewestCubes(lines)
	fmt.Println("Day 01 Part 1 Result: ", partNumSum)
	fmt.Println("Day 01 Part 2 Result: ", gearRatioSum)
}

func sumPartNumbers(lines []string) (int, int) {
	var partNumberSum, gearRatioSum = 0, 0
	var gears = map[Point][]int{}

	h := len(lines)
	w := len(lines[0])

	fmt.Println("The matrix is ", h, "x", w, ".")

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {

			char := lines[row][col]

			if isDigit(char) {
				dcol := col + 1

				for ; dcol < len(lines[row]); dcol++ {
					if !isDigit(lines[row][dcol]) {
						break
					}
				}

				var partNum, _ = strconv.Atoi(lines[row][col:dcol])

				if isPartNum, gearPoint := isPartNumber(lines, col, dcol, row); isPartNum {
					partNumberSum += partNum
					gears[gearPoint] = append(gears[gearPoint], partNum)

				}

				col = dcol
			}
		}
	}

	for _, gearNums := range gears {
		if len(gearNums) == 2 {
			gearRatioSum += gearNums[0] * gearNums[1]
		}
	}

	return partNumberSum, gearRatioSum
}

func isDigit(a byte) bool {
	return a >= '0' && a <= '9'
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPartNumber(schematics []string, x1, x2, y int) (bool, Point) {
	for dy := max(y-1, 0); dy < min(y+2, len(schematics)); dy++ {
		for dx := max(x1-1, 0); dx < min(x2+1, len(schematics[y])); dx++ {
			char := schematics[dy][dx]
			if !(isDigit(char) || char == '.') {
				if char == '*' {
					return true, Point{dx, dy}
				}
				return true, Point{}
			}
		}
	}
	return false, Point{}
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}
