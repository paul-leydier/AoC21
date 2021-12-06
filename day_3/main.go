package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const dataPath = "day_3/data.txt"

// The returned data is represents the columns of input
func readData(path string) ([]string, error) {
	data := make([]string, 12)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			data[i] += line[i : i+1]
		}
	}

	return data, nil
}

// func extractData(data []string) (gamma string, epsilon string) {
// 	for _, datum := range data {
// 		ones := strings.Count(datum, "1")
// 		zeros := strings.Count(datum, "0")
// 		if ones > zeros {
// 			gamma += "1"
// 			epsilon += "0"
// 		} else {
// 			gamma += "0"
// 			epsilon += "1"
// 		}
// 	}
// 	return gamma, epsilon
// }

func extractOxygenGeneratorRating(data []string) string {
	// loop: if more than 1 bit left, keep going
	for i := 0; i < len(data) && len(data[0]) > 1; i++ {
		ones := strings.Count(data[i], "1")
		zeros := strings.Count(data[i], "0")
		if ones >= zeros {
			data = keepOnly(data, '1', i)
		} else {
			data = keepOnly(data, '0', i)
		}
	}
	var oxygenGeneratorRating string
	for _, datum := range data {
		oxygenGeneratorRating += datum
	}
	return oxygenGeneratorRating
}

func extractCO2ScrubberRating(data []string) string {
	// loop: if more than 1 bit left, keep going
	for i := 0; i < len(data) && len(data[0]) > 1; i++ {
		ones := strings.Count(data[i], "1")
		zeros := strings.Count(data[i], "0")
		if ones < zeros {
			data = keepOnly(data, '1', i)
		} else {
			data = keepOnly(data, '0', i)
		}
	}
	var oxygenGeneratorRating string
	for _, datum := range data {
		oxygenGeneratorRating += datum
	}
	return oxygenGeneratorRating
}

func keepOnly(data []string, bit uint8, position int) []string {
	newData := make([]string, len(data))
	for i := 0; i < len(data[0]); i++ {
		if data[position][i] == bit {
			for i2, datum := range data {
				newData[i2] += string(datum[i])
			}
		}
	}
	return newData
}

func main() {
	data, err := readData(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	oxygenGeneratorRatingBits := extractOxygenGeneratorRating(data)
	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorRatingBits, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	CO2ScrubberRatingBits := extractCO2ScrubberRating(data)
	CO2ScrubberRating, err := strconv.ParseInt(CO2ScrubberRatingBits, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Oxygen: %d\nCO2: %d\nLife support: %d", oxygenGeneratorRating, CO2ScrubberRating, oxygenGeneratorRating*CO2ScrubberRating)
}
