package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputAsInts := convertInputToInts(readInput())
	fmt.Println("Answer for day 9, challenge 2 is", getJoltDifferencesMultipliedPartDos(inputAsInts))

}

func getJoltDifferencesMultipliedPartDos(adapters []int) int {
	//count := 0

	adapters = append(adapters, 0)
	// first up. order low to high
	sort.Ints(adapters)
	// pff. must be easier
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	sort.Ints(adapters)

	size := len(adapters)

	solutions := make([]int, size)
	solutions[size-1] = 1

	for i := size - 2; i >= 0; i-- {
		solutions[i] = solutions[i+1]

		// maybe loop
		//check +2 and +3 adapters
		if adapters[i+2] <= adapters[i]+3 && i+2 < size {
			solutions[i] += solutions[i+2]
			//println("got something", adapters[i + 2])
		}
		//} else if adapters[i+3] <= adapters[i] + 3&& i + 3 < size  {
		// nope
		if adapters[i+3] <= adapters[i]+3 && i+3 < size {
			solutions[i] += solutions[i+3]
			//println("got something", adapters[i + 3])
		}
	}
	// answer is max

	max := 0

	for _, option := range solutions {
		if option > max {
			max = option
		}
	}

	return max
}

func convertInputToInts(data []string) []int {
	dataInt := []int{}
	for _, string := range data {
		input, _ := strconv.Atoi(string)
		dataInt = append(dataInt, input)
	}
	return dataInt
}

func readInput() []string {
	var input []string

	f, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
