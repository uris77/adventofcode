package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Command struct {
	Qty int
	Src int
	Dst int
}

func main() {
	one()
	two()
}

func parseCommand(s string) Command {
	var c [3]int
	fmt.Sscanf(s, "move %d from %d to %d", &c[0], &c[1], &c[2])

	return Command{
		Qty: c[0],
		Src: c[1],
		Dst: c[2],
	}
}

// ExtractCol extracts the values for each column into a stack
func ExtractCol(s string, col int) []string {
	splits := strings.Split(s, "\n")
	items := strings.Split(splits[col-1], " ")
	var stack []string
	for i := range items {
		if items[i] == " " {
			stack = append(stack, "")
			continue
		}
		for _, runeValue := range items[i] {
			if unicode.IsLetter(runeValue) {
				stack = append(stack, string(runeValue))
			}
		}

	}
	return stack
}

type stack map[int][]rune

func readInput() stack {
	file, err := os.Open("2022/day5/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	linesCtr := 0
	stacks := make(stack)
	for i := 0; i <= 9; i++ {
		stacks[i] = make([]rune, 0)
	}
	for s.Scan() {
		if linesCtr == 8 {
			return stacks
		}
		line := s.Text()
		for i, box := range line {
			// Catch the 'done' case
			if box == '1' {
				continue
			}
			if (i-1)%4 == 0 && box != ' ' {
				col := ((i - 1) / 4) + 1
				stacks[col] = append([]rune{box}, stacks[col]...)
			}
		}
		linesCtr = linesCtr + 1
	}
	return stacks
}

func readCommands() []Command {
	var commands []Command
	file, err := os.Open("2022/day5/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	linesCtr := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		if linesCtr > 9 {
			line := s.Text()
			c := parseCommand(line)
			commands = append(commands, c)
		}
		linesCtr = linesCtr + 1
	}
	return commands
}

func Move(stacks stack, src, dest, toMove int) stack {
	stacks[dest] = append(stacks[dest], stacks[src][len(stacks[src])-toMove:]...)
	stacks[src] = stacks[src][0 : len(stacks[src])-toMove]
	return stacks
}

func oneMove(stacks stack, from, to int) {
	var box rune
	box, stacks[from] = stacks[from][len(stacks[from])-1], stacks[from][:len(stacks[from])-1]
	stacks[to] = append(stacks[to], box)
}

func Top(stacks stack) string {
	t := ""
	for i := 1; i <= 9; i++ {
		s := stacks[i]
		if len(s) == 0 {
			t = t + " "
		} else {
			t = t + string(s[len(s)-1])
		}
	}
	return t
}

func one() {
	stacks := readInput()
	commands := readCommands()
	for i := range commands {
		for m := 0; m < commands[i].Qty; m++ {
			oneMove(stacks, commands[i].Src, commands[i].Dst)
		}
	}
	top := Top(stacks)
	log.Printf("\n\n:TOP: %s", top)
}

func two() {
	stacks := readInput()
	commands := readCommands()
	for i := range commands {
		Move(stacks, commands[i].Src, commands[i].Dst, commands[i].Qty)
	}
	log.Printf("\nTop: %s", Top(stacks))
}
