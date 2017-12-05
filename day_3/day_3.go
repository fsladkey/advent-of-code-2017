package main

import (
	"fmt"
)

type vector struct {
	x int
	y int
}

var memo = map[string]int{}

var deltas = []vector{
	vector{0, 1},
	vector{1, 0},
	vector{1, 1},
	vector{0, -1},
	vector{-1, 0},
	vector{-1, 1},
	vector{1, -1},
	vector{-1, -1},
}

func (v vector) hash() string {
	return fmt.Sprintf("%d,%d", v.x, v.y)
}

func neighbors(pos vector) []vector {
	neighborVecs := []vector{}
	for _, delta := range deltas {
		neighborVecs = append(neighborVecs, vector{pos.x + delta.x, pos.y + delta.y})
	}
	return neighborVecs
}

func getVal(pos vector) int {
	if val, ok := memo[pos.hash()]; ok {
		return val
	}
	total := 0
	for _, neighbor := range neighbors(pos) {
		total += memo[neighbor.hash()]
	}
	memo[pos.hash()] = total
	return total
}

func getSpiralValueAboveMax(max int) int {
	stepsFromCenter := 1
	pos := vector{0, 0}
	memo[pos.hash()] = 1
	if val := getVal(pos); val > max {
		return val
	}
	for {
		for pos.x < stepsFromCenter {
			pos.x++
			if val := getVal(pos); val > max {
				return val
			}
		}
		for pos.y < stepsFromCenter {
			pos.y++
			if val := getVal(pos); val > max {
				return val
			}
		}
		for pos.x > -stepsFromCenter {
			pos.x--
			if val := getVal(pos); val > max {
				return val
			}
		}
		for pos.y > -stepsFromCenter {
			pos.y--
			if val := getVal(pos); val > max {
				return val
			}
		}
		stepsFromCenter++
	}
}

func getSpiralPosition(input int) (x, y int) {
	stepsFromCenter := 0
	currentNum := 1
	x, y = 0, 0
	if currentNum == input {
		return
	}
	for {
		for x < stepsFromCenter {
			x++
			currentNum++
			if currentNum == input {
				return
			}
		}
		for y < stepsFromCenter {
			y++
			currentNum++
			if currentNum == input {
				return
			}
		}
		for x > -stepsFromCenter {
			x--
			currentNum++
			if currentNum == input {
				return
			}
		}
		for y > -stepsFromCenter {
			y--
			currentNum++
			if currentNum == input {
				return
			}
		}
		stepsFromCenter++
	}
}

func spiralMemory(input int) int {
	x, y := getSpiralPosition(input)
	return x + y
}

func spiralMemory2(input int) int {
	val := getSpiralValueAboveMax(input)
	return val
}

func main() {
	fmt.Println(spiralMemory(361527))
	fmt.Println(spiralMemory2(361527))
}
