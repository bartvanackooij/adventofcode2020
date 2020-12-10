package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputAsInts := convertInputToInts(readInput())

	fmt.Println("Answer for day 9, challenge 1 is", getJoltDifferencesMultiplied(inputAsInts))
}

func getJoltDifferencesMultiplied(adapters []int) int {
	count := 0

	// adapter output; 16,10,15,6,1 etc.
	// adapter input -1,-2,-3 then 16,10,15,6,1 etc

	// +3 van max.  19+3 = 22 in exmaple 1

	// ALLE ADAPTERS 1,4,5,6 .. .. (19+3) gebruikt.

	// 1 (3)
	// 4 (1)
	// 5 (1)
	// 6 (1)
	// 7 (3)
	// 10 (1)
	// 11 (1)
	// 12 (3)
	// 15 (1)
	// 16 (1)
	// 19 (3)
	// 22 (3)

	// 5*3 en 7*1 // answer = 35

	// first up. order low to high
	sort.Ints(adapters)

	// set answer holders
	// get's the first 0 for free as well
	count1 := 1
	// get the last one for free
	count3 := 1

	// loop over each

	for i := 0; i < len(adapters)-1; i++ {
		println(adapters[i])

		// don't forget to check input

		difference := adapters[i+1] - adapters[i]
		switch difference {
		case 0:
			log.Fatal("wrong input, 2 same values: ", adapters[i])
		case 1:
			count1++
		case 2:
			// nothing, but no error neither. I guess, not sure
		case 3:
			count3++
		default:
			// somethings wrong here. bigger than 3
			log.Fatal("wrong input ", adapters[i], " can't reach ", adapters[i+1])
		}
	}

	println("1's: ", count1)
	println("3's: ", count3)

	count = count1 * count3
	return count
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

	f, _ := os.Open("data_test.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
