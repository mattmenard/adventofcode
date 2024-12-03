package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC 2024")
	aoc1203()
}

func aoc1203() {
	lines := getLines("./input2.txt")

	totalSum := 0

	for i := 0; i < len(lines); i++ {
		muls, _ := findMulSequences(lines[i])

		totalSum += executeMultiplicationInstructions(muls)
	}

	fmt.Printf("The total sum of the uncorrupted multiplication instructions is %d.\n", totalSum)

	totalSum = 0

	for i := 0; i < len(lines); i++ {
		muls := extractPatterns(lines[i])

		totalSum += executeInstructions(muls)
	}

	fmt.Printf("The total sum of the uncorrupted multiplication instructions with conditional instructions is %d.\n", totalSum)
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

// findMulSequences searches for sequences that start with "mul(" followed by an integer, a comma,
// another integer, and a closing parenthesis ")".
func findMulSequences(input string) ([]string, error) {
	// Define the regular expression pattern for the desired sequence
	pattern := `mul\(\d+,\d+\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		// Handle errors that occur during regex compilation
		return nil, fmt.Errorf("error compiling regex: %v", err)
	}

	// Find all strings that match the pattern
	matches := re.FindAllString(input, -1)

	return matches, nil
}

func executeMultiplicationInstructions(muls []string) int {
	totalSum := 0

	for i := 0; i < len(muls); i++ {
		mulInst := muls[i]
		mulInst, _ = strings.CutPrefix(mulInst, "mul(")
		mulInst, _ = strings.CutSuffix(mulInst, ")")
		nums := strings.Split(mulInst, ",")

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		prod := num1 * num2

		totalSum += prod

	}

	return totalSum
}

// extractPatterns searches for specific patterns in the input text and returns the matches.
// It looks for the following patterns:
// - Strings that start with "mul(" followed by an integer, a comma, another integer, and a closing parenthesis ")"
// - Strings that start with "do" followed by a set of parentheses "()"
// - Strings that start with "don't" followed by a set of parentheses "()"
func extractPatterns(text string) []string {
	// Define a regular expression pattern that matches the specified strings
	//pattern := `(?i)\b(mul\(\d+,\d+\)|do\(\)|don't\(\))`
	pattern := `mul\(\d+,\d+\)|do\(\)|don't\(\)`

	re := regexp.MustCompile(pattern)

	// Find all matches in the input text
	matches := re.FindAllString(text, -1)

	return matches
}

func executeInstructions(instructions []string) int {
	totalSum := 0
	executeInst := true

	do := 0
	dont := 0
	mul := 0

	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]

		if strings.Contains(inst, "mul") && executeInst {
			mul++
			inst, _ = strings.CutPrefix(inst, "mul(")
			inst, _ = strings.CutSuffix(inst, ")")
			nums := strings.Split(inst, ",")

			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])

			prod := num1 * num2

			fmt.Printf("    Processing MUL instruction %s resulting in %d * %d = %d.\n", inst, num1, num2, prod)

			totalSum += prod
		} else if strings.Contains(inst, "don't") {
			fmt.Println("  ENCOUNTERED DON'T CONDITION, STOPPING FUTURE PROCESSING!")
			dont++
			executeInst = false
		} else if strings.Contains(inst, "do") {
			fmt.Println("  ENCOUNTERED DO CONDITION, RESUMING PROCESSING!")
			do++
			executeInst = true
		} else {
			fmt.Printf("    Processing MUL instruction %s but it is inside a don't block, skipping\n", inst)
			mul++
		}
	}

	//fmt.Printf("do: %d\ndont: %d\nmul: %d\n\n-----------------------\n", do, dont, mul)

	return totalSum
}
