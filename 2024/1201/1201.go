package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC 2024")
	aoc1201()
}

func aoc1201() {
	lines := getLines("./input.txt")

	list1, list2 := getLists(lines)
	distance := calculateTotalDistance(list1, list2)

	fmt.Printf("Total Distance: %f\n", distance)

	score := calculateSimilarityScore(list1, convertSliceToMap(list2))
	fmt.Printf("Similarity Score: %f\n", score)
}

func getLists(lines []string) ([]float64, []float64) {
	list1 := make([]float64, len(lines))
	list2 := make([]float64, len(lines))
	var pos = 0

	for _, line := range lines {
		strs := strings.Split(line, "   ")
		list1[pos], _ = strconv.ParseFloat(strs[0], 64)
		list2[pos], _ = strconv.ParseFloat(strs[1], 64)
		pos++
	}

	sort.Float64s(list1)
	sort.Float64s(list2)

	return list1, list2
}

func calculateTotalDistance(locations1, locations2 []float64) float64 {
	var totalDistance float64

	for i := 0; i < len(locations1); i++ {
		distance := locations1[i] - locations2[i]
		distance = math.Abs(distance)
		//fmt.Printf("Distance: %f - %f = %f\n", locations1[i], locations2[i], distance)
		totalDistance += distance
	}

	return totalDistance
}

func calculateSimilarityScore(list1 []float64, mapOfList2 map[float64]int) float64 {
	var similarityScore float64

	for i := 0; i < len(list1); i++ {
		instances := mapOfList2[list1[i]]
		score := float64(instances) * list1[i]
		similarityScore += score
	}

	return similarityScore
}

// convertSliceToMap takes a slice of float64 and returns a map where the keys are the unique values
// from the slice and the values are the counts of how many times each number appears in the slice.
func convertSliceToMap(slice []float64) map[float64]int {
	result := make(map[float64]int)
	for _, value := range slice {
		result[value]++
	}
	return result
}

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}
