package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSumAndProductOfTwo(t *testing.T) {
	// set test data
	test := readTestInput()

	var input = [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	// solution
	correctSolution := 336

	result := totalTreesHit(input, test)

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
