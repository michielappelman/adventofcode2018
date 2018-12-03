package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/michielappelman/adventofcode2018/pkg/generic"
)

func countSameLetters(id string) []int {
	var sameLetters []int
	counted := make(map[rune]bool)
	for _, r := range id {
		if _, ok := counted[r]; ok {
			continue
		}
		counted[r] = true
		count := strings.Count(id, string(r))
		if !generic.ContainsInt(sameLetters, count) {
			sameLetters = append(sameLetters, count)
		}
	}
	return sameLetters
}

func StarOne(input []string) string {
	counts := make(map[int]int)
	for _, id := range input {
		letterCount := countSameLetters(id)
		for _, c := range letterCount {
			counts[c] += 1
		}
	}
	return strconv.Itoa(counts[2] * counts[3])
}

func offByChars(input string, check string) (int, string) {
	diff := 0
	var sameChars []string
	for i, r := range input {
		if rune(check[i]) != r {
			diff++
		} else {
			sameChars = append(sameChars, string(r))
		}
	}
	return diff, strings.Join(sameChars, "")
}

func StarTwo(input []string) string {
	for _, id := range input {
		for _, check := range input {
			if id == check {
				continue
			}
			if diff, chars := offByChars(id, check); diff == 1 {
				return chars
			}
		}
	}
	return ""
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("1:", StarOne(input))
	fmt.Println("2:", StarTwo(input))
}
