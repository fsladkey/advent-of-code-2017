package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func h(values []int) string {
	valuesText := []string{}

	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	return strings.Join(valuesText, ",")
}

func isRepeat(seen map[string]int, numbers []int) bool {
	if _, ok := seen[h(numbers)]; ok {
		return true
	}
	return false
}

func getMaxIdx(numbers []int) int {
	maxIdx := 0
	for idx, num := range numbers {
		if num > numbers[maxIdx] {
			maxIdx = idx
		}
	}
	return maxIdx
}

func memoryReallocation(filename string) (int, int) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	intChars := strings.Fields(string(contents))
	numbers := []int{}
	for _, char := range intChars {
		numbers = append(numbers, strToInt(char))
	}
	seen := map[string]int{}
	cycles := 0
	for !isRepeat(seen, numbers) {
		seen[h(numbers)] = cycles

		maxIdx := getMaxIdx(numbers)
		maxVal := numbers[maxIdx]
		numbers[maxIdx] = 0
		for idx := 0; idx < maxVal; idx++ {
			numbers[(maxIdx+1+idx)%len(numbers)]++
		}
		cycles++
	}
	return cycles, cycles - seen[h(numbers)]
}

func main() {
	// fmt.Println(memoryReallocation("./test_input.txt"))
	cycles, diff := memoryReallocation("./input.txt")
	fmt.Println(cycles)
	fmt.Println(diff)
}
