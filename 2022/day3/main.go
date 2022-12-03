package main

import (
	"bufio"
	"log"
	"os"
)

var charValues = []rune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	in := readInput()
	one(in)
	two(in)
}

func one(in [][]rune) {
	// Split each array of runes into half
	var halvedRunes [][][]rune
	for i, _ := range in {
		half := HalveRunes(in[i])
		halvedRunes = append(halvedRunes, half)
	}
	// Find the intersection runes
	var common []rune
	for i, _ := range halvedRunes {
		c := Intersection(halvedRunes[i])
		if len(c) != 0 {
			common = append(common, c...)
		}
	}
	// Sum them up
	log.Printf("\nSum: %d", Sum(common))
}

func two(in [][]rune) {
	// Split into groups of 3
	grouped := [][][]rune{{}}
	var currentGroup [][]rune
	for i, _ := range in {
		isEndOfGroup := (i+1)%3 == 0
		currentGroup = append(currentGroup, in[i])
		if isEndOfGroup {
			grouped = append(grouped, currentGroup)
			currentGroup = [][]rune{}
		}
	}
	//log.Printf("Grouped: %c", grouped)
	// Find common item in each group
	var labels []rune
	for i, _ := range grouped {
		if len(grouped[i]) > 0 {
			common := Intersection3(grouped[i])
			labels = append(labels, common...)
		}
	}
	log.Printf("\nAll Priorities: %d", Sum(labels))
}

func readInput() [][]rune {
	file, err := os.Open("2022/day3/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	var in [][]rune
	s := bufio.NewScanner(file)
	for s.Scan() {
		txt := s.Text()
		// Convert strings to array of runes
		var r []rune
		for _, runeValue := range txt {
			r = append(r, runeValue)
		}
		in = append(in, r)
	}
	return in
}

func HalveRunes(in []rune) [][]rune {
	halfSize := len(in) / 2
	out := [][]rune{{}, {}}
	for idx, runeValue := range in {
		if idx+1 <= halfSize {
			out[0] = append(out[0], runeValue)
		} else {
			out[1] = append(out[1], runeValue)
		}
	}
	return out
}

func Intersection(in [][]rune) []rune {
	var out []rune
	// Keep track of items that have already been found in the other rucksack
	// If rucksack A has abradr and B has dfaer, then r should only appear once
	// in the list of common items
	var currentDuplicates []rune
	for i, runeValue := range in[0] {
		if Index(runeValue, in[1]) > -1 && Index(runeValue, currentDuplicates) == -1 {
			out = append(out, runeValue)
			currentDuplicates = append(currentDuplicates, runeValue)
		}
		if i+1 == len(in[0]) {
			currentDuplicates = []rune{}
		}
	}
	return out
}
func Intersection3(in [][]rune) []rune {
	var out []rune
	// Keep track of items that have already been found in the other rucksack
	// If rucksack A has abradr and B has dfaer, then r should only appear once
	// in the list of common items
	var currentDuplicates []rune
	for i, runeValue := range in[0] {
		if Index(runeValue, in[1]) > -1 &&
			Index(runeValue, in[2]) > -1 &&
			Index(runeValue, currentDuplicates) == -1 {
			out = append(out, runeValue)
			currentDuplicates = append(currentDuplicates, runeValue)
		}
		if i+1 == len(in[0]) {
			currentDuplicates = []rune{}
		}
	}
	return out
}

func Sum(in []rune) int {
	total := 0
	for _, runeValue := range in {
		if v := Index(runeValue, charValues); v > 0 {
			total = total + v
		}
	}

	return total
}
func Index(r rune, in []rune) int {
	for i, runeValue := range in {
		if r == runeValue {
			return i
		}
	}
	return -1
}
