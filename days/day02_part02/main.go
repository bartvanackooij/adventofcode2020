package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Answer for day 2, challenge 2 is", passwordPhilosopher(readInput()))
}

//func passwordPhilosopher(data []string) int {
func passwordPhilosopher(data []string) int {
	count := 0

	// loop over all passwordLines
	for _, firstInput := range data {
		collection := strings.Fields(firstInput)

		// policy is e.g. 1-3, 2-9
		policy := collection[0]

		// policyMin = lowest, policyMax = highest
		policyMin, _ := strconv.Atoi(strings.Split(policy, "-")[0])
		policyMax, _ := strconv.Atoi(strings.Split(policy, "-")[1])

		// policySubject is e.g. a, b, c
		policySubject := collection[1]
		//remove the ":"
		policySubject = strings.Split(policySubject, ":")[0]

		// password is e.g. abcde, cdefg, ccccccc
		password := collection[2]

		// boolean to see if they hit;
		firstHit := string(password[policyMin-1]) == policySubject
		secondHit := string(password[policyMax-1]) == policySubject

		// check both booleans
		if firstHit != secondHit {
			count++
		}
	}
	return count
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
