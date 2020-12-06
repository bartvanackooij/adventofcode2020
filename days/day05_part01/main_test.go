package main

import (
	"log"
	"testing"
)

func TestNarrowDownScope(t *testing.T) {
	testEstimationInput := [][]int{
		{0, 127},
		{0, 63},
		{32, 63},
		{32, 47},
		{40, 47},
		{44, 45},
		{0, 127},
	}
	testLetterInput := []string{
		"F",
		"B",
		"F",
		"B",
		"B",
		"F",
		"B",
	}
	correctSolutions := [][]int{
		{0, 63},
		{32, 63},
		{32, 47},
		{40, 47},
		{44, 47},
		{44, 44},
		{64, 127},
	}

	for i := 0; i < len(testEstimationInput); i++ {
		estimation := testEstimationInput[i]          // {0,127}
		letter := testLetterInput[i]                  // F
		correctSolution := correctSolutions[i]        // {0,63}
		result := narrowDownScope(estimation, letter) // []int uit.

		if result[0] != correctSolution[0] {
			log.Fatal("e-", result[0], correctSolution[0])
		}

		if result[1] != correctSolution[1] {
			log.Fatal("e+", result[1], correctSolution[1])
		}
	}
}
