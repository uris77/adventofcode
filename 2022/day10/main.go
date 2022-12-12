package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var signalStrengths = map[int]int{
	20:  0,
	60:  0,
	100: 0,
	140: 0,
	180: 0,
	220: 0,
}

func main() {
	onev1()
	input := parseInput()
	log.Printf("\nOne v2: %d\n", onev2(input))
	log.Printf("Two: \n%s", two(input))
}

func onev2(input []int) int {
	total := 0
	for _, x := range []int{20, 60, 100, 140, 180, 220} {
		total = total + (input[x-1] * x)
	}
	return total
}

func parseInput() []int {
	file, err := os.Open("2022/day10/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	var cycles []int
	register := 1
	cycles = append(cycles, register)
	s := bufio.NewScanner(file)
	for s.Scan() {
		txt := s.Text()
		switch txt {
		case "noop":
			cycles = append(cycles, register)
		default:
			v := 0
			fmt.Sscanf(txt, "addx %d", &v)
			cycles = append(cycles, register)
			register = register + v
			cycles = append(cycles, register)
		}
	}
	return cycles
}

func onev1() {
	file, err := os.Open("2022/day10/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	cycles := 0
	register := 1
	s := bufio.NewScanner(file)
	for s.Scan() {
		v := 0
		txt := s.Text()
		switch txt {
		case "noop":
			cycles = cycles + 1
			signalStrength(cycles, register)
		default:
			fmt.Sscanf(txt, "addx %d", &v)
			cycles = cycles + 1
			signalStrength(cycles, register)
			cycles = cycles + 1
			register = register + v
			signalStrength(cycles, register)
		}
	}
	all := 0
	for i := 20; i <= 220; i = i + 40 {
		all = all + signalStrengths[i]

	}
	log.Printf("One: %d", all)
}

func signalStrength(cycle, register int) {
	if cycle != 19 && cycle != 59 && cycle != 99 && cycle != 139 && cycle != 179 && cycle != 219 {
		return
	}
	signalStrengths[cycle+1] = (cycle + 1) * register
}

// answer should be EFGERURE
func two(cycles []int) string {
	var renderedLines []string

	currentScreen := ""
	for c := 0; c < len(cycles); c++ {
		pixel := len(currentScreen)
		currentPosition := cycles[c]
		if (pixel == currentPosition) || (pixel == currentPosition-1) || (pixel == currentPosition+1) {
			currentScreen += "#"
		} else {
			currentScreen += "."
		}
		if len(currentScreen) == 40 {
			renderedLines = append(renderedLines, currentScreen)
			//Clear screen so we can start the next line
			currentScreen = ""
		}
	}
	return strings.Join(renderedLines, "\r\n")
}
