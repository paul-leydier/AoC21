package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const dataPath = "day_1/data.csv"

func readIntsFromCsv(path string) ([]int64, error) {
	var data []int64
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, number)
	}

	return data, nil
}

func countIncreases(measures []int64) int {
	count := 0
	for i := 0; i < len(measures)-1; i++ {
		if measures[i] < measures[i+1] {
			count++
		}
	}
	return count
}

func sum(window []int64) int {
	var sum int64
	for _, x := range window {
		sum += x
	}
	return int(sum)
}

func countWindowIncreases(measures []int64) int {
	count := 0
	const windowsSize = 3
	truncHalfWindowSize := windowsSize / 2 // integer division
	previousWindow := sum(measures[:windowsSize])
	for i := windowsSize - truncHalfWindowSize; i < len(measures)-truncHalfWindowSize; i++ {
		currentWindow := sum(measures[i-truncHalfWindowSize : i+truncHalfWindowSize+1])
		if currentWindow > previousWindow {
			count++
		}
		previousWindow = currentWindow
	}
	return count
}

func main() {
	measures, err := readIntsFromCsv(dataPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Part 1 - Number of depth increases: %d\n", countIncreases(measures))

	fmt.Printf("Part 2 - Number of depth increases: %d\n", countWindowIncreases(measures))
}
