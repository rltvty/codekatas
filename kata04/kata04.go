package kata04

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Weather() int {
	file, err := os.Open("./weather.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	skipRows := 2

	for skipRows > 0 && scanner.Scan() {
		skipRows--
	}

	maxDiff := 0
	maxDiffDay := 0

	for scanner.Scan() {
		columns, err := getColumns(scanner.Text(), []int{0, 1, 2}, "*")
		if err == nil && columns[0] != "mo" {
			maxTemp, err := strconv.Atoi(columns[1])
			if err != nil {
				fmt.Printf("Could not convert %v to int\n", columns[1])
			}
			minTemp, err := strconv.Atoi(columns[2])
			if err != nil {
				fmt.Printf("Could not convert %v to int\n", columns[2])
			}
			if maxTemp-minTemp > maxDiff {
				maxDiff = maxTemp - minTemp
				maxDiffDay, err = strconv.Atoi(columns[0])
				if err != nil {
					fmt.Printf("Could not convert %v to int\n", columns[0])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return maxDiffDay
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
