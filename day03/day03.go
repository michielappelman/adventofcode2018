package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/michielappelman/adventofcode2018/pkg/generic"
)

type Square struct {
	col, row int
}

func StarOne(input []string) string {
	fabric := make(map[Square][]int)
	pattern := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	for _, claim := range input {
		matches := pattern.FindAllStringSubmatch(claim, 1)
		ints, _ := generic.StringsToInts(matches[0][1:])
		for x := ints[1]; x < ints[1]+ints[3]; x++ {
			for y := ints[2]; y < ints[2]+ints[4]; y++ {
				fabric[Square{x, y}] = append(fabric[Square{x, y}], ints[0])
			}
		}
	}
	double := 0
	for _, s := range fabric {
		if len(s) > 1 {
			double++
		}
	}
	return strconv.Itoa(double)
}

func StarTwo(input []string) string {
	fabric := make(map[Square][]int)
	pattern := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	ids := make(map[int]bool)
	for _, claim := range input {
		matches := pattern.FindAllStringSubmatch(claim, 1)
		ints, _ := generic.StringsToInts(matches[0][1:])
		ids[ints[0]] = true
		for x := ints[1]; x < ints[1]+ints[3]; x++ {
			for y := ints[2]; y < ints[2]+ints[4]; y++ {
				fabric[Square{x, y}] = append(fabric[Square{x, y}], ints[0])
			}
		}
	}
	for _, s := range fabric {
		if len(s) > 1 {
			for _, c := range s {
				delete(ids, c)
			}
		}
	}
	var left int
	for k, _ := range ids {
		left = k
	}
	return strconv.Itoa(left)
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
