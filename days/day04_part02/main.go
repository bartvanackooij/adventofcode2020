package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Answer for day 4, challenge 2 is", validatePassport(readInput()))
}

var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validatePassport(data []string) int {
	count := 0
	passport := ""
	var testData = []string{}

	for _, passportInput := range data {
		if len(passportInput) > 0 {
			var valid = true
			for _, field := range strings.Split(passportInput, " ") {
				if valid {
					for _, mandatoryField := range mandatoryFields {
						if strings.HasPrefix(field, mandatoryField) {
							switch mandatoryField {
							case "byr":
								valid = checkByr(field)
								break
							case "iyr":
								valid = checkIyr(field)
								break
							case "eyr":
								valid = checkEyr(field)
								break
							case "hgt":
								valid = checkHgt(field)
								break
							case "hcl":
								valid = checkHcl(field)
								break
							case "ecl":
								valid = checkEcl(field)
								break
							case "pid":
								valid = checkPid(field)
								break
							}
						}
					}
				}
			}

			if valid {
				strings.Join(testData, passportInput)
				passport += passportInput
			}
		} else {
			if checkMandatoryFields(passport) {
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
	input = append(input, "")
	return input
}

func removeKey(keyvalue string) string {
	// in byr:1937
	return strings.Split(keyvalue, ":")[1]
	// out 1937
}

// byr:1937
func checkByr(byr string) bool {
	birthyear, _ := strconv.Atoi(removeKey(byr))
	return birthyear >= 1920 && birthyear <= 2003
}

// iyr:2003
func checkIyr(iyr string) bool {
	issueYear, _ := strconv.Atoi(removeKey(iyr))
	return issueYear >= 2010 && issueYear <= 2020
}

// eyr:2023
func checkEyr(eyr string) bool {
	expirationYear, _ := strconv.Atoi(removeKey(eyr))
	return expirationYear >= 2020 && expirationYear <= 2030
}

// hgt:190cm / 60in
func checkHgt(hgt string) bool {
	if strings.HasSuffix(hgt, "cm") {
		heightStringCm := strings.Split(hgt, "cm")[0]
		heightInCm, _ := strconv.Atoi(removeKey(heightStringCm))
		return heightInCm >= 150 && heightInCm <= 193
	} else if strings.HasSuffix(hgt, "in") {
		heightStringInch := strings.Split(hgt, "in")[0]
		heightInInch, _ := strconv.Atoi(removeKey(heightStringInch))
		return heightInInch >= 59 && heightInInch <= 76
	} else {
		return false
	}
}

// hcl:#ae17e1
func checkHcl(hcl string) bool {
	prefix := "#"
	colorCode := removeKey(hcl)
	// colorcode left: e.g. #ae17e1

	if strings.HasPrefix(colorCode, prefix) {
		//	remove prefix
		code := strings.SplitAfter(colorCode, prefix)[1]
		// golang regex
		// a-f lower
		// A-F capital
		// 0-9 numerics
		// {6 digits}

		match, _ := regexp.MatchString("[a-fA-F0-9]{6}", code)
		return match
	}
	return false
}

// ecl:amb ;blu ;brn ;gry ;grn ;hzl ;oth
func checkEcl(ecl string) bool {
	colorCode := removeKey(ecl)
	// colorcode left: e.g. amb

	for validColorIndex := 0; validColorIndex < len(eyeColors); validColorIndex++ {
		if colorCode == eyeColors[validColorIndex] {
			return true
		}
	}
	return false
}

// pid:012345678
func checkPid(pid string) bool {
	personId := removeKey(pid)
	// colorcode left: e.g. 00000001 ; 012345678

	// golang regex
	// 0-9 numerics
	// {6 digits}

	if len(personId) == 9 {
		match, _ := regexp.MatchString("[0-9]{9}", personId)
		return match
	} else {
		return false
	}
}
