package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	totalFuel := 0

	for _, input := range readInput() {
		mass, _ := strconv.Atoi(input)
		totalFuel += fuelPerModule(mass)
	}

	fmt.Println("Fuel: ", totalFuel)
}

func fuelPerModule(mass int) int {
	return mass/3 - 2
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
