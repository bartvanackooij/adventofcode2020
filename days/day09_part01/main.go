package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	convertInputToInts(readInput())

	fmt.Println("Answer for day 9, challenge 1 is", getErrorEncryptionAnswer(convertInputToInts(readInput())))
}

var preamble = 25

func hasSumCalculation(sum int, data []int) bool {
	println("Starting here; ")

	for i := 0; i < len(data)-1; i++ {
		first := data[i]
		for second := i + 1; second < len(data); second++ {
			//println("first: \t\t", first)
			//println("data[second]: \t", data[second])
			if first+data[second] == sum {
				println(first, "&", data[second])
				return true
			}
		}
	}
	return false
}

func getErrorEncryptionAnswer(data []int) int {

	// loop over 5-20
	for i := preamble; i < len(data); i++ {
		//println(data[i-preamble])

		sum := data[i]
		println("SUM: ", sum)

		// find within 0-5
		startIndex := i - preamble
		endIndex := i
		println("Looking in between ", data[startIndex], " - ", data[endIndex-1])

		if !hasSumCalculation(sum, data[startIndex:endIndex+1]) {
			println("the end")
			return data[i]
		}
	}
	return 0
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
