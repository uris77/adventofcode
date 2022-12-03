package main

import (
	"bufio"
	"golang.org/x/exp/constraints"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("2022/day1/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var currentCalories = 0
	var groupedCalories []int
	for scanner.Scan() {
		txt := scanner.Text()
		// An empty string indicates the end of an elve's calories. So we reset
		// the total calories calculated.
		if txt == "" {
			groupedCalories = append(groupedCalories, currentCalories)
			currentCalories = 0
			continue
		}
		if calories, err := strconv.Atoi(txt); err == nil {
			currentCalories = currentCalories + calories
		}
	}
	// The last calories calculates is appended to the list. This is skipped when we initially
	// scan through the calories because we exit early.
	if currentCalories > 0 {
		groupedCalories = append(groupedCalories, currentCalories)
	}
	sortSlice(groupedCalories)
	log.Printf("Top 3: %d %d %d\n", groupedCalories[0], groupedCalories[1], groupedCalories[2])
	log.Printf("Total top 3: %d\n", groupedCalories[0]+groupedCalories[1]+groupedCalories[2])
	log.Print("\n\n\n\n\n")
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}
