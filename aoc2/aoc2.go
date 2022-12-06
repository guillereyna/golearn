package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// part 1
var shape_scores = map[string]int{
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, // scissors
}

// points if opponent[player]
var rock = map[string]int{
	"X": 3,
	"Y": 6,
	"Z": 0,
}

var paper = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var scissors = map[string]int{
	"X": 6,
	"Y": 0,
	"Z": 3,
}

var round_scores = map[string]map[string]int{
	"A": rock,
	"B": paper,
	"C": scissors,
}

// part 2
var expected_scores = map[string]int{
	"X": 0, // losw
	"Y": 3, // draw
	"Z": 6, // win
}

var vsrock = map[string]int{
	"X": 3,
	"Y": 1,
	"Z": 2,
}

var vspaper = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var vsscissors = map[string]int{
	"X": 2,
	"Y": 3,
	"Z": 1,
}

var vsround_scores = map[string]map[string]int{
	"A": vsrock,
	"B": vspaper,
	"C": vsscissors,
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	totalp2 := 0
	for scanner.Scan() {
		round := scanner.Text()
		shapes := strings.Fields(round)
		opponent := shapes[0]
		player := shapes[1]

		total += shape_scores[player] + round_scores[opponent][player]

		totalp2 += expected_scores[player] + vsround_scores[opponent][player]
	}
	fmt.Println("part1:", total)
	fmt.Println("part2:", totalp2)
}
