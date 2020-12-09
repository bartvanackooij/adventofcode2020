package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var listOfBags = convertInputToBags(readInput())
var listOfColorsForAnswer = []string{"shiny gold"}
var count = 0

func main() {

	bagName := "shiny gold"
	//startBag := new(Bag)
	//startBag.color = bagName

	println(findIndexOfBag(bagName, listOfBags))

	fmt.Println("Answer for day 7, challenge 2 is", getBagsInsideBagsInside(listOfBags[findIndexOfBag(bagName, listOfBags)]))
}

func findIndexOfBag(bagcolor string, listOfBags []*Bag) int {
	for k, v := range listOfBags {
		if bagcolor == v.color {
			return k
		}
	}
	return -1
}

func getBagsInsideBagsInside(bag *Bag) int {
	count := 0
	for insideOf, quantity := range bag.contains {
		count += quantity + (quantity * getBagsInsideBagsInside(listOfBags[findIndexOfBag(insideOf, listOfBags)]))
	}
	return count
}

type Bag struct {
	color    string
	contains map[string]int
	//canBeContainedIn []string
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
