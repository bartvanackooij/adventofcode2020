package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSumAndProductOfThree(t *testing.T) {
	// set test data
	test := readTestInput()

	// solution
	correctSolution := 241861950

	result := sumAndProductOfThree(test)

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
