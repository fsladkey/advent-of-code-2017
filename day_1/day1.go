package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func getNextIdx(input string, idx int) int {
	return (idx + 1) % len(input)
}

func getCircleIdx(input string, idx int) int {
	return (idx + len(input)/2) % len(input)
}

func inverseCaptcha(input string, getNextIdx func(input string, idx int) int) int {
	count := 0
	for idx := range input {
		if input[idx] == input[getNextIdx(input, idx)] {
			char := string(input[idx])
			num, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			count += num
		}
	}
	return count
}

func inverseCaptchaPt1(input string) int {
	return inverseCaptcha(input, getNextIdx)
}

func inverseCaptchaPt2(input string) int {
	return inverseCaptcha(input, getCircleIdx)
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	str := string(input)
	fmt.Println(inverseCaptchaPt1(str))
	fmt.Println(inverseCaptchaPt2(str))
}
