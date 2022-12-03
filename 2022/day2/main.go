package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	win  = 6
	draw = 3
	loss = 0
)

type Hand struct {
	Name   string
	Points int
}

var rock = Hand{
	Name:   "rock",
	Points: 1,
}

var paper = Hand{
	Name:   "paper",
	Points: 2,
}

var scissors = Hand{
	Name:   "scissors",
	Points: 3,
}

var hands = map[string]Hand{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var instructions = map[string]int{
	"Y": draw,
	"X": loss,
	"Z": win,
}

type Outcomes struct {
	Win  string
	Loss string
	Draw string
}

var permutations = map[string]Outcomes{
	"A": {
		Win:  "Y",
		Draw: "X",
		Loss: "Z",
	},
	"B": {
		Win:  "Z",
		Draw: "Y",
		Loss: "X",
	},
	"C": {
		Win:  "X",
		Draw: "Z",
		Loss: "Y",
	},
}

func PlayHand(oppo, you string) int {
	oppHand := hands[oppo]
	youHand := hands[you]

	if oppHand.Points == youHand.Points {
		return draw
	}

	if oppHand.Name == "rock" && youHand.Name == "scissors" {
		return loss
	}

	if oppHand.Name == "scissors" && youHand.Name == "rock" {
		return win
	}

	if oppHand.Points > youHand.Points {
		return loss
	}
	return win
}

func main() {
	file, err := os.Open("2022/day2/input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	var input [][]string
	s := bufio.NewScanner(file)
	for s.Scan() {
		txt := s.Text()
		turn := strings.Split(txt, " ")
		input = append(input, turn)
	}

	log.Printf("\nOne: %d", partOne(input))
	log.Printf("\nTwo: %d", partTwo(input))
}

func partOne(input [][]string) int {
	var pts []int
	for i, _ := range input {
		turn := input[i]
		roundPoints := PlayHand(turn[0], turn[1])
		pts = append(pts, roundPoints+hands[turn[1]].Points)
	}
	return sumPoints(pts)
}

func partTwo(in [][]string) int {
	var pts []int
	for i, _ := range in {
		turn := in[i]
		instruction := instructions[turn[1]]
		hand := ""
		switch instruction {
		case win:
			hand = permutations[turn[0]].Win
		case loss:
			hand = permutations[turn[0]].Loss
		case draw:
			hand = permutations[turn[0]].Draw
		}
		roundPoints := PlayHand(turn[0], hand)
		pts = append(pts, roundPoints+hands[hand].Points)
	}

	return sumPoints(pts)
}

func sumPoints(p []int) int {
	acc := 0
	for i, _ := range p {
		acc = p[i] + acc
	}
	return acc
}
