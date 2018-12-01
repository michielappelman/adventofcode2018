package main

import (
	"bufio"
	"fmt"
	"os"
)

func StarOne(input []string) string {
	return input[0]
}

func StarTwo(input []string) string {
	return input[0]
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
