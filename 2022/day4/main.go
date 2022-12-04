package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	fmt.Printf("One: %d\n", one(input))
	fmt.Printf("Two: %d", two(input))
}

func readInput() [][][]int {
	file, err := os.Open("2022/day4/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	var in [][][]int
	s := bufio.NewScanner(file)
	for s.Scan() {
		txt := s.Text()
		sections := strings.Split(txt, ",")
		sectionOne := ExpandSections(sections[0])
		sectionTwo := ExpandSections(sections[1])
		in = append(in, [][]int{sectionOne, sectionTwo})
	}
	return in
}

func ExpandSections(in string) []int {
	bounds := strings.Split(in, "-")
	start, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])
	return createRange(start, end)
}

func createRange(start, end int) []int {
	s := make([]int, end-start+1)

	for i := range s {
		s[i] = start + i
	}
	return s
}

// IsContainedIn returns true if first is contained in second.
// This is determined using the algorithm: If second[start] >= first[start] && second[end] <= first[end] then
// first is contained in second
func IsContainedIn(first, second []int) bool {
	return second[0] >= first[0] && second[len(second)-1] <= first[len(first)-1]
}

func one(in [][][]int) int {
	cnt := 0

	for i := range in {
		if IsContainedIn(in[i][0], in[i][1]) || IsContainedIn(in[i][1], in[i][0]) {
			cnt = cnt + 1
		}
	}
	return cnt
}

func two(in [][][]int) int {
	cnt := 0

	for i := range in {
		if Overlaps(in[i][0], in[i][1]) {
			cnt = cnt + 1
		}
	}

	return cnt
}

// Overlaps determines if there are overlaps between first and second.
func Overlaps(first, second []int) bool {
	greaterThanLowerBound := first[0] > second[0] && first[len(first)-1] > second[0]
	greaterThanUpperBound := first[0] > second[len(second)-1] && first[len(first)-1] > second[len(second)-1]
	lesserThanLowerBound := first[0] < second[0] && first[len(first)-1] < second[0]
	lesserThanUpperBound := first[0] < second[len(second)-1] && first[len(first)-1] < second[len(second)-1]

	return !((greaterThanLowerBound && greaterThanUpperBound) || (lesserThanLowerBound && lesserThanUpperBound))

}
