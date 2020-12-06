package main

import (
	"bufio"
	"os"
	"testing"
)

func TestValidatePassport(t *testing.T) {
	// solution
	correctSolution := 2

	// set test data
	test := readTestInput()
	result := validatePassport(test)

	if result != correctSolution {
		t.Fatalf("Calculated is %d, expected %d", result, correctSolution)
	}
}

func readTestInput() []string {
	var input []string

	f, _ := os.Open("data_test.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
