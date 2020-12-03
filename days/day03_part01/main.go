package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Answer for day 3, challenge 1 is", hittingTrees(readInput(), 3, 1))
}

// x = 3
// y = 1
//
// y,x
// 0,0
// 1,3
// 2,6
// 3,9
// 4,12 (12%11) 4,1
// 5,15 (15%11) 4,4

func hittingTrees(data []string, x int, y int) int {
	treesHit := 0
	positionY := 0
	positionX := 0

	// loop over all lines every iteration y = plussed
	for positionY < len(data) {
		treeLine := data[positionY]

		if string(treeLine[positionX%len(treeLine)]) == "#" {
			treesHit++
		}

		positionX += x
		positionY += y
	}
	return treesHit
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
