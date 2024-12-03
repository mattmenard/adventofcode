package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC 2024")
	aoc1202()
}

func aoc1202() {
	lines := getLines("./input.txt")

	safeReports := processReports1(lines)
	fmt.Printf("Part1: There are %d safe reports.\n", safeReports)

	safeReports = processReports2(lines)
	fmt.Printf("Part2: There are %d safe reports.\n", safeReports)
}

func processReports1(lines []string) int {
	safeReports := 0

	for _, report := range lines {
		isValid := validateSequence_Part1(report)

		if isValid {
			safeReports++
		}
	}

	return safeReports
}

func processReports2(lines []string) int {
	safeReports := 0

	for _, report := range lines {
		isValid := validateSequence_Part2(report)

		//fmt.Printf("The sequence [%s] is valid: %t\n", report, isValid)

		if isValid {
			safeReports++
		}
	}

	return safeReports
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

// isMonotonic checks if the given slice of integers is either entirely non-increasing or non-decreasing.
func isMonotonic(numbers []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			decreasing = false
		} else if numbers[i] > numbers[i+1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

// checkDifferences checks if the absolute difference between any two adjacent numbers is at least 1 and at most 3.
func checkDifferences(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i] - numbers[i+1]
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

// validateSequence takes a space-delimited string of numbers and checks if the sequence is either
// monotonic increasing or decreasing and if any two adjacent numbers differ by at least 1 and at most 3.
func validateSequence_Part1(input string) bool {
	// Split the input string by spaces
	strNumbers := strings.Split(input, " ")
	var numbers []int

	// Convert strings to integers
	for _, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return false
		}
		numbers = append(numbers, num)
	}

	// Check if the sequence is monotonic and differences are within the specified range
	return isMonotonic(numbers) && checkDifferences(numbers)
}

//=============================================================================
//----------------------------------- PART 2 ----------------------------------
//=============================================================================

func validateSequence_Part2(input string) bool {
	isValid := false

	// Split the input string by spaces
	strNumbers := strings.Split(input, " ")
	var numbers []int

	// Convert strings to integers
	for _, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			return false
		}
		numbers = append(numbers, num)
	}

	isMonotonic := isMonotonic(numbers)
	isGradualChange := checkDifferences(numbers)

	if !isMonotonic || !isGradualChange {
		isValid = applyProblemDampener(numbers)
	} else {
		isValid = true
	}

	return isValid
}

func applyProblemDampener(nums []int) bool {
	isValid := false

	for i := 0; i < len(nums); i++ {
		subNums := make([]int, 0)
		subNums = append(subNums, nums[:i]...)
		subNums = append(subNums, nums[i+1:]...)

		isMonotonic := isMonotonic(subNums)
		isGradualChange := checkDifferences(subNums)

		if isMonotonic && isGradualChange {
			isValid = true
			break
		}
	}

	return isValid
}
