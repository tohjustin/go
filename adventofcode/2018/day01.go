package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func findSum(offsets []int) int {
	sum := 0
	for _, val := range offsets {
		sum += val
	}

	return sum
}

func findDuplicate(offsets []int) int {
	frequencyMap, sum := make(map[int]int), 0
	for i := 0; ; i++ {
		if i >= len(offsets) {
			i = 0
		}

		frequencyMap[sum] = frequencyMap[sum] + 1
		sum += offsets[i]

		if frequencyMap[sum] > 0 {
			return sum
		}
	}
}

func getFrequencyOffsets() []int {
	result := []int{}

	file, err := os.Open(*inputFile)
	if err != nil {
		return result
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		if offset, err := strconv.Atoi(line); err == nil {
			result = append(result, offset)
		}
	}

	return result
}

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	offsets := getFrequencyOffsets()

	resultA := findSum(offsets)
	fmt.Printf("Part A Result: %d\n", resultA)
	resultB := findDuplicate(offsets)
	fmt.Printf("Part B Result: %d\n", resultB)
}
