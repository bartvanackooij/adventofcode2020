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
	correctSolution := 7

	result := hittingTrees(test, 3, 1)

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
