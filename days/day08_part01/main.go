package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Answer for day 8, challenge 1 is", getAccumulatorCode(readInput()))
}

func getAccumulatorCode(data []string) int {
	accumulatorScore := 0
	position := 1
	indexVisited := []int{position}
	infiniteLoop := false
	count := 0

	for !infiniteLoop {
		command := data[position-1]
		//println("count is ", count)
		println("command is ", command)

		if strings.HasPrefix(command, "nop ") {

			// up the position so we can check next step.
			position++
			// check if we've been on the new position before
			if checkForInfiniteLoop(position, indexVisited) {
				break
			}
			// if it is, we'll break. otherwise. add it to indexvisited.
			indexVisited = append(indexVisited, position)
			println("nop: new position is ", position)
		}
		if strings.HasPrefix(command, "acc ") {
			params := strings.Split(command, "acc ")[1]
			if strings.HasPrefix(params, "+") {
				amount, _ := strconv.Atoi(strings.Split(params, "+")[1])
				// positive
				//println("+oud", accumulatorScore)
				accumulatorScore += amount
				println("ACC NEW", accumulatorScore)
			} else {
				amount, _ := strconv.Atoi(strings.Split(params, "-")[1])
				//println("-oud", accumulatorScore)
				accumulatorScore -= amount
				println("ACC NEW", accumulatorScore)
			}
			position++
			if checkForInfiniteLoop(position, indexVisited) {
				break
			}
			indexVisited = append(indexVisited, position)
			println("acc: new position is ", position)
		}
		if strings.HasPrefix(command, "jmp ") {
			params := strings.Split(command, "jmp ")[1]
			if strings.HasPrefix(params, "+") {
				amount, _ := strconv.Atoi(strings.Split(params, "+")[1])
				// positive
				position += amount
			} else {
				amount, _ := strconv.Atoi(strings.Split(params, "-")[1])
				position -= amount
			}
			if checkForInfiniteLoop(position, indexVisited) {
				break
			}
			indexVisited = append(indexVisited, position)
		}
		println()
		count++
	}
	return accumulatorScore
}

func checkForInfiniteLoop(position int, visited []int) bool {
	println("checking position", position, " size is", len(visited))
	infiniteLoop := false
	for _, visitedPosition := range visited {
		//println("VP", visitedPosition)
		//println("P", position)

		if visitedPosition == position {
			// been there. infinite loop
			println("Stop de persen!!")
			infiniteLoop = true
			return infiniteLoop
		}
	}
	return infiniteLoop
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
