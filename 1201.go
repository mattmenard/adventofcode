package main

import (
	"fmt"
	"os"
	"strings"
)

var numbers = map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func aoc1201() {
	lines := getLines("./input.txt")

	var result = getCalibrationSum(lines, false)
	var result2 = getCalibrationSum(lines, true)
	fmt.Println("Day 01 Part 1 Result: ", result)
	fmt.Println("Day 01 Part 2 Result: ", result2)
}

func getCalibrationSum(lines []string, part2 bool) int {
	var calibrationVal int

	for _, line := range lines {
		fmt.Println(line)
		var first, second int
	firstNum:
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = int(line[i] - '0')
				break
			} else if part2 {
				for k := range numbers {
					if strings.HasPrefix(line[i:], k) {
						first = numbers[k]
						break firstNum
					}
				}
			}
		}
	secondNum:
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				second = int(line[i] - '0')
				break
			} else if part2 {
				for k := range numbers {
					if strings.HasPrefix(line[i:], k) {
						second = numbers[k]
						break secondNum
					}
				}
			}
		}

		calibrationVal += first*10 + second
	}
	return calibrationVal
}
