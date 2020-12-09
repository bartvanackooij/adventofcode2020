package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Answer for day 8, challenge 2 is", changeInput(readInput()))
}

func getAccumulatorCode(data []string) int {
	accumulatorScore := 0
	position := 1
	indexVisited := []int{position}
	infiniteLoop := false
	notInfiniteLoopButDone := false
	count := 0

	for !infiniteLoop {
		command := data[position-1]
		if strings.HasPrefix(command, "nop ") {
			position++
			if checkForInfiniteLoop(position, indexVisited) {
				infiniteLoop = true
				break
			}
			indexVisited = append(indexVisited, position)
		}
		if strings.HasPrefix(command, "acc ") {
			params := strings.Split(command, "acc ")[1]
			if strings.HasPrefix(params, "+") {
				amount, _ := strconv.Atoi(strings.Split(params, "+")[1])
				accumulatorScore += amount
			} else {
				amount, _ := strconv.Atoi(strings.Split(params, "-")[1])
				accumulatorScore -= amount
			}
			position++
			if checkForInfiniteLoop(position, indexVisited) {
				infiniteLoop = true
				break
			}
			indexVisited = append(indexVisited, position)
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
				infiniteLoop = true
				break
			}
			indexVisited = append(indexVisited, position)
		}
		if len(command) == 0 {
			// the end !
			notInfiniteLoopButDone = true
			break
		}
		count++
	}

	if infiniteLoop {
		accumulatorScore = 0
	}

	if notInfiniteLoopButDone {
		return accumulatorScore
	}

	return accumulatorScore
}

func checkForInfiniteLoop(position int, visited []int) bool {
	infiniteLoop := false
	for _, visitedPosition := range visited {
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

	f, _ := os.Open("data_test.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	input = append(input, "")
	return input
}

func changeInput(data []string) int {
	result := 0

	indexList := []int{}
	count := 0
	for _, line := range data {

		if strings.HasPrefix(line, "jmp ") {
			indexList = append(indexList, count)
		}
		count++
	}

	for _, indexOfJMP := range indexList {
		updatedString := strings.Replace(data[indexOfJMP], "jmp", "nop", -1)
		data[indexOfJMP] = updatedString

		if getAccumulatorCode(data) != 0 {
			result = getAccumulatorCode(data)
			break
		} else {
			// set back
			returnString := strings.Replace(data[indexOfJMP], "nop", "jmp", -1)
			data[indexOfJMP] = returnString
		}
	}
	return result
}
