package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func sortCharacters(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func lineIsValid1(line string) bool {
	seen := map[string]bool{}
	words := strings.Fields(line)
	for _, word := range words {
		if _, ok := seen[word]; ok {
			return false
		}
		seen[word] = true
	}
	return true
}

func lineIsValid2(line string) bool {
	seen := map[string]bool{}
	words := strings.Fields(line)
	for _, word := range words {
		for key := range seen {
			if strings.Compare(sortCharacters(word), sortCharacters(key)) == 0 {
				return false
			}
		}
		seen[word] = true
	}
	return true
}

func highEntropyPassphrases(filename string, lineIsValid func(line string) bool) int {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	count := 0
	for _, line := range lines {
		if lineIsValid(line) {
			count++
		}
	}
	return count
}

func highEntropyPassphrases1(filename string) int {
	return highEntropyPassphrases(filename, lineIsValid1)
}

func highEntropyPassphrases2(filename string) int {
	return highEntropyPassphrases(filename, lineIsValid2)
}

func main() {
	fmt.Println(highEntropyPassphrases1("./input.txt"))
	fmt.Println(highEntropyPassphrases2("./input.txt"))
}
