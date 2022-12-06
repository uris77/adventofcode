package main

import (
	"log"
	"os"
)

func main() {
	rawString, err := os.ReadFile("2022/day6/input.txt")
	if err != nil {
		panic("Could not find input file")
	}
	window := slidingWindow(4, []rune(string(rawString)))
	one(window)
	secondWindow := slidingWindow(14, []rune(string(rawString)))
	two(secondWindow)
}

func one(window [][]rune) {
	cnt := 0
	var marker []rune
	for i := range window {
		d := hasDiffs(window[i])
		if d == true && cnt == 0 {
			cnt = i
			marker = window[i]
		}
	}
	log.Printf("\ncnt: %d ; marker: %c ; start: %d", cnt, marker, cnt+4)
}

func two(window [][]rune) {
	//log.Printf("window: %c", window)
	cnt := 0
	var som []rune
	for i := range window {
		d := hasDiffs(window[i])
		if d == true && cnt == 0 {
			cnt = i
			som = window[i]
		}
	}

	log.Printf("\ncnt: %d; som: %c; start: %d", cnt, som, cnt+14)
}

func hasDiffs(in []rune) bool {
	yes := true
	prev := make(map[rune]bool, 0)
	for i := 0; i < len(in); i++ {
		if prev[in[i]] == true {
			return false
		} else {
			prev[in[i]] = true
		}
	}

	return yes
}

func slidingWindow(size int, input []rune) [][]rune {
	// returns the input slice as the first element
	if len(input) <= size {
		return [][]rune{input}
	}

	// allocate slice at the precise size we need
	r := make([][]rune, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}
