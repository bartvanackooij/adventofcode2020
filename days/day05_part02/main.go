package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	fmt.Println("Answer for day 5, challenge 2 is", getAirplaneSeatId(readInput()))
}

func getAirplaneSeatId(data []string) int {
	// loop over each seat
	var seatIds = []int{}
	var mySeat int
	for _, seatCode := range data {
		seatId := calculateSeatId(seatCode)
		seatIds = append(seatIds, seatId)
	}
	sort.Ints(seatIds)

	// find the gap
	for i := 0; i < len(seatIds)-1; i++ {
		if seatIds[i]+1 != seatIds[i+1] {
			// myseat is in between
			mySeat = (seatIds[i+1] + seatIds[i]) / 2
		}
	}
	return mySeat
}

func calculateSeatId(seatCode string) int {
	var row int
	var column int

	seatRowEstimate := []int{0, 127}
	// retrieve row
	for i := 0; i < 7; i++ {
		sliceWithRowCodes := strings.Split(seatCode, "")
		seatRowEstimate = narrowDownScope(seatRowEstimate, sliceWithRowCodes[i])
	}
	if seatRowEstimate[0] == seatRowEstimate[1] {
		row = seatRowEstimate[0]
	}
	// retrieve column
	seatColumnEstimate := []int{0, 7}
	for i := 7; i < 10; i++ {
		sliceWithColumnCodes := strings.Split(seatCode, "")
		seatColumnEstimate = narrowDownScope(seatColumnEstimate, sliceWithColumnCodes[i])
	}
	if seatColumnEstimate[0] == seatColumnEstimate[1] {
		column = seatColumnEstimate[0]
	}

	seatId := row*8 + column
	return seatId
}

func narrowDownScope(estimation []int, letter string) []int {
	if letter == "F" || letter == "L" {
		// lower half, cut high 32/63
		estimation[1] = estimation[1] - ((estimation[1]-estimation[0])+1)/2
	} else {
		if estimation[1]%2 == 1 {
		}
		estimation[0] = estimation[0] + (((estimation[1] + 1) - estimation[0]) / 2)
	}
	return estimation
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
