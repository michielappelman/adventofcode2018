package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const numOfWorkers = 2
const additionalSeconds = 0

func isPossible(parents map[string][]string, instr string, been string) bool {
	n := 0
	for _, p := range parents[instr] {
		if strings.Contains(been, p) {
			n++
		}
	}
	if n == len(parents[instr]) {
		return true
	}
	return false
}

func availableInstructions(parents map[string][]string, been string) []string {
	var possible []string
	for c, _ := range parents {
		if isPossible(parents, c, been) && !strings.Contains(been, c) {
			possible = append(possible, c)
		}
	}
	sort.Strings(possible)
	return possible
}

func availableInstructionsTwo(parents map[string][]string, been string, done string) []string {
	var possible []string
	for c, _ := range parents {
		if isPossible(parents, c, done) && !strings.Contains(been, c) {
			possible = append(possible, c)
		}
	}
	sort.Strings(possible)
	return possible
}

func findRoots(parents map[string][]string) []string {
	var roots []string
	for c, p := range parents {
		if len(p) == 0 {
			roots = append(roots, c)
		}
	}
	return roots
}

func parseParents(input []string) map[string][]string {
	parents := make(map[string][]string)
	pattern := regexp.MustCompile(`^Step (\w) .* step (\w) .*\.$`)
	for _, line := range input {
		matches := pattern.FindAllStringSubmatch(line, 1)
		parents[matches[0][2]] = append(parents[matches[0][2]], matches[0][1])
		if _, ok := parents[matches[0][1]]; !ok {
			parents[matches[0][1]] = []string{}
		}
	}
	return parents
}

func StarOne(input []string) string {
	parents := parseParents(input)

	been := findRoots(parents)[0]
	for next := availableInstructions(parents, been); len(next) > 0; next = availableInstructions(parents, been) {
		been = fmt.Sprintf("%s%s", been, next[0])
	}
	return been
}

func getDuration(instr string) int {
	return additionalSeconds + (int(rune(instr[0])) - 64)
}

func StarTwo(input []string) string {
	freeWorkers := numOfWorkers
	parents := parseParents(input)
	second := 0
	doneOn := make(map[int]string)
	next := findRoots(parents)
	been := strings.Join(next, "")
	done := ""

	for len(been) != len(done) || len(next) > 0 {
		for _, instr := range next {
			if freeWorkers > 0 {
				if !strings.Contains(been, instr) {
					been = fmt.Sprintf("%s%s", been, instr)
				}
				doneOn[second+getDuration(instr)] = fmt.Sprintf("%s%s", doneOn[second+getDuration(instr)], instr)
				freeWorkers--
			}
		}
		second++
		freeWorkers += len(doneOn[second])
		done = fmt.Sprintf("%s%s", done, doneOn[second])
		next = availableInstructionsTwo(parents, been, done)
	}

	return strconv.Itoa(second)
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
