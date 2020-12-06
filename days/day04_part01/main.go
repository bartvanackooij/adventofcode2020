package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Answer for day 4, challenge 1 is", validatePassport(readInput()))
}

var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func validatePassport(data []string) int {
	count := 0
	// create empty passport string array for storing data
	passport := ""

	// break data in lines / short bits blabla
	for _, passportInput := range data {
		if len(passportInput) > 0 {

			// passportInput it not empty. need to append to existing passport.
			passport += passportInput

		} else {
			// data is empty, means new line in data. check passport and create empty one.
			if checkMandatoryFields(passport) {
				println(passport, " is valid. ")
				count++
			}
			passport = ""
		}
	}
	return count
}

func checkMandatoryFields(passport string) bool {
	isValid := true
	for _, mandatoryField := range mandatoryFields {
		if !strings.Contains(passport, mandatoryField) {
			return false
		}
	}
	return isValid
}

func readInput() []string {
	var input []string

	f, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	// because of the loop, append a new line at the end for this challenge
	input = append(input, "")
	return input
}
