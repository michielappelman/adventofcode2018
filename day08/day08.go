package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/michielappelman/adventofcode2018/pkg/generic"
)

func getNodeLengthSum(node []int) (int, int) {
	numOfChildNodes := node[0]
	numOfMetadata := node[1]

	length := 2
	sum := 0

	for i := 0; i < numOfChildNodes; i++ {
		childLength, childSum := getNodeLengthSum(node[length:])
		length += childLength
		sum += childSum
	}

	sum += generic.Sum(node[length : length+numOfMetadata])
	length += numOfMetadata

	return length, sum
}

func StarOne(input []string) string {
	fields := strings.Fields(input[0])
	tree, _ := generic.StringsToInts(fields)
	_, sum := getNodeLengthSum(tree)

	return strconv.Itoa(sum)
}

func getNodeLengthValue(node []int) (int, int) {
	numOfChildNodes := node[0]
	numOfMetadata := node[1]

	length := 2
	if numOfChildNodes == 0 {
		return length + numOfMetadata, generic.Sum(node[length : length+numOfMetadata])
	}

	childValues := make(map[int]int)

	for i := 0; i < numOfChildNodes; i++ {
		childLength, childValue := getNodeLengthValue(node[length:])
		length += childLength
		childValues[i+1] = childValue
	}

	value := 0
	metadata := node[length : length+numOfMetadata]
	for _, m := range metadata {
		if c, ok := childValues[m]; ok {
			value += c
		}
	}

	length += numOfMetadata

	return length, value
}

func StarTwo(input []string) string {
	fields := strings.Fields(input[0])
	tree, _ := generic.StringsToInts(fields)
	_, value := getNodeLengthValue(tree)

	return strconv.Itoa(value)
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
