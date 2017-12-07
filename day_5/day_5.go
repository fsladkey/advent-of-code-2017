package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func postJumpAdjustment1(offset int) int {
	return offset + 1
}

func postJumpAdjustment2(offset int) int {
	if offset >= 3 {
		return offset - 1
	}
	return postJumpAdjustment1(offset)
}

func twistyTrampolines(filename string, postJumpAdjustment func(int) int) int {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	numbers := []int{}
	for _, line := range strings.Split(string(contents), "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	jumps := 0
	idx := 0
	for idx < len(numbers) {
		jumps++
		offset := numbers[idx]

		numbers[idx] = postJumpAdjustment(offset)
		idx += offset
	}
	return jumps
}

func twistyTrampolines1(filename string) int {
	return twistyTrampolines(filename, postJumpAdjustment1)
}

func twistyTrampolines2(filename string) int {
	return twistyTrampolines(filename, postJumpAdjustment2)
}

func main() {
	// fmt.Println(twistyTrampolines("./test_input.txt"))
	fmt.Println(twistyTrampolines1("./input.txt"))
	fmt.Println(twistyTrampolines2("./input.txt"))
}
