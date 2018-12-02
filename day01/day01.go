package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/michielappelman/adventofcode2018/pkg/generic"
)

func StarOne(input []string) string {
	ints, err := generic.StringsToInts(input)
	if err != nil {
		log.Fatal("could not convert to ints")
	}
	return strconv.Itoa(generic.Sum(ints))
}

func StarTwo(input []string) string {
	startingFrequency := 0
	ints, err := generic.StringsToInts(input)
	if err != nil {
		log.Fatal("could not convert to ints")
	}
	seen := make(map[int]struct{})
	seen[startingFrequency] = struct{}{}
	for {
		for _, a := range ints {
			startingFrequency += a
			if _, ok := seen[startingFrequency]; ok {
				return strconv.Itoa(startingFrequency)
			}
			seen[startingFrequency] = struct{}{}
		}
	}
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
