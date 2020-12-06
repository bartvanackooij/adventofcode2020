package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Answer for day 6, challenge 2 is", questionCount(readInput()))
}

func questionCount(data []string) int {
	count := 0
	totalGroupCount := 0
	answerAsString := []string{}
	numberOfPeople := 0
	for _, groupAnswer := range data {

		if len(groupAnswer) > 0 {
			numberOfPeople++

			splitAnswers := strings.Split(groupAnswer, "")
			for _, answer := range splitAnswers {
				answerAsString = append(answerAsString, answer)
			}
		} else {
			groupCount := calcualteGroupCount2(answerAsString, numberOfPeople)
			totalGroupCount = totalGroupCount + groupCount
			answerAsString = []string{}
			numberOfPeople = 0

			count++
		}
	}
	return totalGroupCount
}

func calcualteGroupCount2(groupAnswer []string, people int) int {
	newCount := 0
	// pretend to loop over alphabet
	var groupCount = 1
	for i := 1; i < 27; i++ {
		letterCount := 0
		for _, answer := range groupAnswer {
			if toChar(i) == answer {
				letterCount++
				if letterCount == people {
					newCount++
				}
			}
		}
	}
	groupCount++
	return newCount
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
