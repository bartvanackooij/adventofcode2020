package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSumAndProductOfTwo(t *testing.T) {
	// set test data
	test := readTestInput()

	// solution
	correctSolution := 514579

	result := sumAndProductOfTwo(test)

	if result != correctSolution {
		t.Fatalf("Calculated sum is %d, expected %d", result, correctSolution)
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
