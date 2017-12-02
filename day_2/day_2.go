package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func minMaxDifference(row []string) int {
	firstVal, err := strconv.Atoi(row[0])
	if err != nil {
		panic(err)
	}
	largest := firstVal
	smallest := firstVal
	for _, char := range row {
		num, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		if num > largest {
			largest = num
		} else if num < smallest {
			smallest = num
		}
	}
	return largest - smallest
}

func evenlyDivisibleDifference(row []string) int {
	for idx := 0; idx < len(row); idx++ {
		num1, err := strconv.Atoi(row[idx])
		if err != nil {
			panic(err)
		}
		for idy := 1; idy < len(row); idy++ {
			num2, err := strconv.Atoi(row[(idx+idy)%len(row)])
			if err != nil {
				panic(err)
			}
			if num1%num2 == 0 {
				return num1 / num2
			}
		}
	}
	return 0
}

func corruptionChecksum(grid []string, getSumOfDifference func(row []string) int) int {
	sumOfDifferences := 0
	for _, line := range grid {
		row := strings.Fields(line)
		sumOfDifferences += getSumOfDifference(row)
	}
	return sumOfDifferences
}

func corruptionChecksum1(grid []string) int {
	return corruptionChecksum(grid, minMaxDifference)
}

func corruptionChecksum2(grid []string) int {
	return corruptionChecksum(grid, evenlyDivisibleDifference)
}

func main() {
	grid, err := readLines("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(corruptionChecksum1(grid))
	fmt.Println(corruptionChecksum2(grid))
}
