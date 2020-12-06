package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Answer for day 6, challenge 1 is", questionCount(readInput()))
}

func questionCount(data []string) int {
	count := 0
	totalGroupCount := 0
	answerAsString := []string{}

	for _, groupAnswer := range data {
		if len(groupAnswer) > 0 {
			//	 append to data
			splitAnswers := strings.Split(groupAnswer, "")
			for _, answer := range splitAnswers {
				answerAsString = append(answerAsString, answer)
			}

		} else {
			// loop over a-z. if hit, go to next, and. up groupcount.
			groupCount := calcualteGroupCount(answerAsString)
			totalGroupCount = totalGroupCount + groupCount
			answerAsString = []string{}
			count++
		}
	}
	return totalGroupCount
}

func calcualteGroupCount(groupAnswer []string) int {
	count := 0
	// pretend to loop over alphabet
	for i := 1; i < 27; i++ {
		for _, answer := range groupAnswer {
			if toChar(i) == answer {
				// found the letter in the answer.
				// up count and break this letter.
				count++
				break
			}
		}
	}
	return count
}

func toChar(i int) string {
	return string('a' - 1 + i)
}

func readInput() []string {
	var input []string

	f, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	input = append(input, "")
	return input
}
