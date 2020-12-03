package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = [][]int{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func main() {
	fmt.Println("Answer for day 3, challenge 2 is", totalTreesHit(input, readInput()))
}

func totalTreesHit(routes [][]int, data []string) int {
	totalHit := 1

	for _, route := range routes {
		totalHit = totalHit * hittingTrees(data, route[0], route[1])
	}

	return totalHit
}

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
