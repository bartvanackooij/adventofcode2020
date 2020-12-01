package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sumGoal = 2020

func main() {
	fmt.Println("Answer for day 1, challenge 1 is", sumAndProductOfTwo(readInput()))
}

func sumAndProductOfTwo(data []string) int {
	var answer int
	// loop over all numbers
	for _, firstInput := range data {
		firstNumber, _ := strconv.Atoi(firstInput)
		// for each number, loop again and show the sum.
		for _, secondInput := range readInput() {
			secondNumber, _ := strconv.Atoi(secondInput)
			// check the goal
			if firstNumber+secondNumber == sumGoal {
				answer = firstNumber * secondNumber
			}
		}
	}
	return answer
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
