package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func collapsePolymer(polymer []rune) int {
	prevLen := 0
	polymerLength := len(polymer)
	for ; polymerLength != prevLen; polymerLength = len(polymer) {
		prevLen = len(polymer)
		for i, c := range polymer {
			if i+1 < len(polymer) {
				if polymer[i+1] == c+32 || polymer[i+1] == c-32 {
					polymer = append(polymer[:i], polymer[i+2:]...)
				}
			}
		}
	}
	return polymerLength
}

func StarOne(input []string) string {
	var polymer []rune
	for _, c := range input[0] {
		polymer = append(polymer, c)
	}
	return strconv.Itoa(collapsePolymer(polymer))
}

func StarTwo(input []string) string {
	var polymer []rune
	typeCollapse := make(map[rune]int)
	for _, c := range input[0] {
		polymer = append(polymer, c)
		typeCollapse[unicode.ToLower(c)] = 0
	}
	for c, _ := range typeCollapse {
		var newPolymer []rune
		for _, r := range polymer {
			if r == c || r == unicode.ToUpper(c) {
				continue
			}
			newPolymer = append(newPolymer, r)
		}
		typeCollapse[c] = collapsePolymer(newPolymer)
	}

	var lowest int
	for _, lowest = range typeCollapse {
		break
	}
	for _, l := range typeCollapse {
		if l < lowest {
			lowest = l
		}
	}
	return strconv.Itoa(lowest)
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
