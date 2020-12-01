package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sumGoal = 2020

func main() {
	fmt.Println("Answer for day 1, challenge 2 is", sumAndProductOfThree(readInput()))
}

func sumAndProductOfThree(data []string) int {
	var answer int
	// loop over all numbers
	for _, firstInput := range data {
		firstNumber, _ := strconv.Atoi(firstInput)
		// for each number, loop again and show the sum.
		for _, secondInput := range data {
			secondNumber, _ := strconv.Atoi(secondInput)
			// and again!
			for _, thirdInput := range data {
				thirdNumber, _ := strconv.Atoi(thirdInput)
				// and check if matches the goal.
				if (firstNumber + secondNumber + thirdNumber) == sumGoal {
					answer = firstNumber * secondNumber * thirdNumber
					break
				}
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
