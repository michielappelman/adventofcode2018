package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

var TEST_INPUT_FILE = "test_input.txt"

func TestSolutions(t *testing.T) {
	content, err := ioutil.ReadFile(TEST_INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	days := strings.Split(string(content), "\n///\n")
	dayOneTests := strings.Split(days[0], "\n===\n")
	for i, test := range dayOneTests {
		testDetails := strings.Split(test, "\n---\n")
		input := strings.Split(testDetails[0], "\n")
		got := StarOne(input)
		if got != testDetails[1] {
			t.Errorf("for %s got %s, want %s", input, got, testDetails[1])
		} else {
			fmt.Printf("Test %d succeeded, got %s\n", i+1, got)
		}
	}

	if len(days) >= 2 {
		dayTwoTests := strings.Split(days[1], "\n===\n")
		for i, test := range dayTwoTests {
			testDetails := strings.Split(test, "\n---\n")
			input := strings.Split(testDetails[0], "\n")
			got := StarTwo(input)
			if got != testDetails[1] {
				t.Errorf("for %s got %s, want %s", input, got, testDetails[1])
			} else {
				fmt.Printf("Test %d succeeded, got %s\n", i+1, got)
			}
		}
	}
}
