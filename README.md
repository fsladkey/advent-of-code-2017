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