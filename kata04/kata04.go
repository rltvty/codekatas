package kata04

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func weather() string {
	maxDiff := 0
	maxDiffDay := ""

	for i, columns := range getDataFromFile("./weather.dat", []int{0, 1, 2}) {
		if i == 0 {
			continue //skip header row
		}
		if columns[0] != "mo" {
			maxTemp, err := strconv.Atoi(columns[1])
			if err != nil {
				fmt.Printf("Could not convert %v to int\n", columns[1])
				continue
			}
			minTemp, err := strconv.Atoi(columns[2])
			if err != nil {
				fmt.Printf("Could not convert %v to int\n", columns[2])
				continue
			}
			if diff := maxTemp - minTemp; diff > maxDiff {
				maxDiff = diff
				maxDiffDay = columns[0]
			}
		}
	}
	return maxDiffDay
}

func team() string {
	minDiff := 100000.0
	minDiffTeam := ""

	for i, columns := range getDataFromFile("./football.dat", []int{1, 6, 8}) {
		if i == 0 {
			continue //skip header row
		}
		goalsFor, err := strconv.Atoi(columns[1])
		if err != nil {
			fmt.Printf("Could not convert %v to int\n", columns[1])
			continue
		}
		goalsAgainst, err := strconv.Atoi(columns[2])
		if err != nil {
			fmt.Printf("Could not convert %v to int\n", columns[2])
			continue
		}
		if diff := math.Abs(float64(goalsFor) - float64(goalsAgainst)); diff < minDiff {
			minDiff = diff
			minDiffTeam = columns[0]
		}
	}
	return minDiffTeam
}

func getDataFromFile(filename string, columns []int) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columns, err := getColumns(scanner.Text(), columns, "*")
		if err == nil {
			output = append(output, columns)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func getColumns(row string, columns []int, cutset string) ([]string, error) {
	cols := strings.Fields(row)
	output := make([]string, len(columns))
	for i := 0; i < len(columns); i++ {
		if columns[i] >= len(cols) {
			return nil, errors.New("Specified column does not exist in row")
		}
		output[i] = strings.Trim(cols[columns[i]], cutset)
	}
	return output, nil
}
