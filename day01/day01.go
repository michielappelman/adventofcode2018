package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func StarOne(input []string) int {
	startingFrequency := 0
	for _, a := range input {
		if a[0] == 43 {
			digit, _ := strconv.Atoi(string(a[1:]))
			startingFrequency = startingFrequency + digit
		} else {
			digit, _ := strconv.Atoi(string(a[1:]))
			startingFrequency = startingFrequency - digit
		}
	}
	return startingFrequency
}

func StarTwo(input []string) int {
	startingFrequency := 0
	seen := make(map[int]struct{})
	seen[startingFrequency] = struct{}{}
	for {
		for _, a := range input {
			if a[0] == 43 {
				digit, _ := strconv.Atoi(string(a[1:]))
				startingFrequency = startingFrequency + digit
				if _, ok := seen[startingFrequency]; ok {
					return startingFrequency
				}
				seen[startingFrequency] = struct{}{}
			} else {
				digit, _ := strconv.Atoi(string(a[1:]))
				startingFrequency = startingFrequency - digit
				if _, ok := seen[startingFrequency]; ok {
					return startingFrequency
				}
				seen[startingFrequency] = struct{}{}
			}
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
