package main

import (
	"bufio"
	"os"
	"testing"
)

func readTestInput() []string {
	var input []string

	f, _ := os.Open("data_test.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func TestRemoveKey(t *testing.T) {
	testData := []string{
		"eyr:1972",
		"cid:100",
		"hcl:#18171d",
		"ecl:amb",
		"hgt:170",
		"pid:186cm",
		"iyr:2018",
		"byr:1926"}
	correctSolution := []string{
		"1972",
		"100",
		"#18171d",
		"amb",
		"170",
		"186cm",
		"2018",
		"1926"}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := removeKey(temp)
		if result != correctSolution[test] {
			t.Fatalf("Calculated is %s, expected %s", result, correctSolution[test])
		}
		test++
	}
	println("TestRemoveKey is success")
}

func TestCheckByr(t *testing.T) {
	testData := []string{
		"byr:2002",
		"byr:2003"}

	correctSolutions := []bool{
		true,
		false}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := checkByr(temp)
		if result != correctSolutions[test] {
			t.Fatalf("Calculated is %t, expected %t", result, correctSolutions[test])
		}
		test++
	}
	println("TestCheckByr is success")
}

func TestCheckIyr(t *testing.T) {
	testData := []string{
		"iyr:2017",
		"iyr:2013",
		"iyr:2007",
		"iyr:2021"}

	correctSolutions := []bool{
		true,
		true,
		false,
		false}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := checkIyr(temp)
		if result != correctSolutions[test] {
			t.Fatalf("Calculated is %t, expected %t", result, correctSolutions[test])
		}
		test++
	}
	println("TestCheckIyr is success")
}

func TestCheckEyr(t *testing.T) {
	testData := []string{
		"eyr:2020",
		"eyr:2023",
		"eyr:2018",
		"eyr:2045"}

	correctSolutions := []bool{
		true,
		true,
		false,
		false}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := checkEyr(temp)
		if result != correctSolutions[test] {
			t.Fatalf("Calculated is %t, expected %t", result, correctSolutions[test])
		}
		test++
	}
	println("TestCheckEyr is success")
}

func TestCheckHgt(t *testing.T) {
	testData := []string{
		"hgt:60in",
		"hgt:190cm",
		"hgt:190in",
		"hgt:190"}

	correctSolutions := []bool{
		true,
		true,
		false,
		false}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := checkHgt(temp)
		if result != correctSolutions[test] {
			t.Fatalf("Calculated is %t, expected %t", result, correctSolutions[test])
		}
		test++
	}
	println("TestCheckHgt is success")
}

func TestCheckHcl(t *testing.T) {
	testData := []string{
		"hcl:#123abc",
		"hcl:#123abz",
		"hcl:123abc"}

	correctSolutions := []bool{
		true,
		false,
		false}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := checkHcl(temp)
		if result != correctSolutions[test] {
			t.Fatalf("Calculated is %t, expected %t", result, correctSolutions[test])
		}
		test++
	}
	println("TestCheckHcl is success")
}

func TestCheckEcl(t *testing.T) {
	testData := []string{
		"ecl:brn",
		"ecl:gry",
		"ecl:oth",
		"ecl:gnr",
		"ecl:isa"}

	correctSolutions := []bool{
		true,
		true,
		true,
		false,
		false}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := checkEcl(temp)
		if result != correctSolutions[test] {
			t.Fatalf("Calculated is %t, expected %t", result, correctSolutions[test])
		}
	}
	println("TestCheckEHcl is success")
}

func TestCheckPid(t *testing.T) {
	testData := []string{
		"pid:000000001",
		"pid:005000001",
		"pid:000000001",
		"pid:0123a56789",
		"pid:0123456789",
		"pid:0123456789"}

	correctSolutions := []bool{
		true,
		true,
		true,
		false,
		false,
		false}

	for test := 0; test < len(testData); test++ {
		temp := testData[test]
		result := checkPid(temp)
		if result != correctSolutions[test] {
			t.Fatalf("Calculated is %t, expected %t", result, correctSolutions[test])
		}
	}
	println("TestCheckPid is success")
}
