package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// const for checking if value already exists
const t = true

func main() {

	listOfBags := convertInputToBags(readInput())
	fmt.Println("Answer for day 7, challenge 1 is", getBagsThatCanContainX(listOfBags, "shiny gold"))

}

var listOfColorsForAnswer = []string{"shiny gold"}

func getBagsThatCanContainX(listOfBags []*Bag, colorToFind string) int {

	// create map to see if item exists

	for _, bag := range listOfBags {
		if bag.contains[colorToFind] != 0 {
			println("GOT ONE! ", bag.color)

			// need to add to listOfColorsForAnswer, if it's not already there.

			exists := false
			for _, itemInList := range listOfColorsForAnswer {
				println("comparing:", itemInList, bag.color)
				if itemInList == bag.color {
					exists = true
				}
			}
			if !exists {
				println("-------------------")
				println(bag.color, "ADDING ", bag.color)
				println("-------------------")
				listOfColorsForAnswer = append(listOfColorsForAnswer, bag.color)
			}
			//
			////count++
			newColor := bag.color
			getBagsThatCanContainX(listOfBags, newColor)
		}
	}
	// minus the original one.
	return len(listOfColorsForAnswer) - 1
}

type Bag struct {
	color            string
	contains         map[string]int
	canBeContainedIn []string
}

func convertInputToBags(data []string) []*Bag {

	var listOfBags = []*Bag{}

	for _, line := range data {
		// let's get the different lines from the example ;

		// each line is a bag.
		bag := new(Bag)
		//bag = new(Bag)

		// first 2 is name of the bag .
		dataSplit := strings.Split(line, " bags contain ")

		color := dataSplit[0]
		//println("kleur:", color)
		bag.color = color

		containsString := dataSplit[1]
		// remove the ones that have none;

		if containsString != "no other bags." {
			// contains other bags, might be multiple
			// split on bag, bags, bags. bag.
			// regex to use ; bag(s?)(,|\.)
			contains := regexp.MustCompile("bag(s?)(, |. )").Split(containsString, -1)
			containMap := make(map[string]int)
			for _, contain := range contains {
				// now here. it's like this ;
				//1 bright white
				//2 muted yellow
				//1 bright white bag.
				containSingleString := strings.Split(contain, " ")

				// as a map;

				// set color it contains as key
				colorName := containSingleString[1] + " " + containSingleString[2]
				colorQuantity, _ := strconv.Atoi(containSingleString[0])
				//println(colorName, colorQuantity)
				containMap[colorName] = colorQuantity
				bag.contains = containMap
			}
		}
		listOfBags = append(listOfBags, bag)
	}
	return listOfBags
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
