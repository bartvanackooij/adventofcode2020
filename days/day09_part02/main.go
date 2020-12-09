package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	convertInputToInts(readInput())

	//answer := getErrorEncryptionAnswer(convertInputToInts(readInput()))
	answer := 258585477
	//answer := 127
	//fmt.Println("Answer for day 9, challenge 1 is", getErrorEncryptionAnswer(convertInputToInts(readInput())))
	fmt.Println("Answer for day 9, challenge 2 is", findSumOfAnswerOrSomething(answer, convertInputToInts(readInput())))
}

var preamble = 25

func findSumOfAnswerOrSomething(answer int, data []int) int {

	// start with 0 .
	// add [0+1]
	// add [0+2]
	//
	// if > answer break and next
	// if == answer win

	size := len(data)

	for i := 0; i < size; i++ {
		// at least 2 numbers
		for j := i; j < size; j++ {
			//listToHoldOptions := []int{}
			//startIndex := i
			//endIndex := j
			listToHoldOptions := data[i:j]
			temp := 0

			//println("Looking for Answer : ", answer)
			//println("Starting with first number", data[i])
			//println("Size of listToHoldOptions ", len(listToHoldOptions))

			for k := 0; k < len(listToHoldOptions) && temp < answer; k++ {
				//println("adding : ", listToHoldOptions[k])
				temp += listToHoldOptions[k]
				//println("tempSum is ", temp)
				//println("Looking for ", answer)
			}

			if temp == answer {
				println("Got something ", listToHoldOptions[0])
				println("and  ", listToHoldOptions[len(listToHoldOptions)-1])

				// now find max here;

				min := 9999999999
				max := 0

				for _, option := range listToHoldOptions {
					println(option)
					if option > max {
						max = option
					}
					if option < min {
						min = option
					}
				}

				return min + max
			}
		}
	}
	return 0
}

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
