### Day 1

```go
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
```

### Day 2
```go
func minMaxDifference(row []string) int {
    firstVal := strToInt(row[0])
    largest := firstVal
    smallest := firstVal
    for _, char := range row {
        num := strToInt(char)
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
        num1 := strToInt(row[idx])
        for idy := 1; idy < len(row); idy++ {
            num2 := strToInt(row[(idx+idy)%len(row)])
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
```

### Day 3
```go
// ew
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
```