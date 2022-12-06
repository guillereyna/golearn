package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(character rune) int {
	if character < 97 {
		return int(character - 38)
	}
	return int(character - 96)
}

func fill_counter(rucksack string, counter []bool) {
	for _, item := range rucksack {
		counter[priority(item)] = true
	}
}

func part1(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		var firstcounter, secondcounter [53]bool

		rucksack := scanner.Text()
		half := len(rucksack) / 2
		first, second := rucksack[0:half], rucksack[half:]

		fill_counter(first, firstcounter[:])
		fill_counter(second, secondcounter[:])

		for index := range firstcounter {
			if firstcounter[index] && secondcounter[index] {
				total += index
			}
		}
	}

	return total
}

func part2(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		var firstcounter, secondcounter, thirdcounter [53]bool

		fill_counter(first, firstcounter[:])
		fill_counter(second, secondcounter[:])
		fill_counter(third, thirdcounter[:])

		for index := range firstcounter {
			if firstcounter[index] && secondcounter[index] && thirdcounter[index] {
				total += index
			}
		}
	}

	return total
}

func main() {
	file, _ := os.Open("input.txt")
	fmt.Println("part1:", part1(bufio.NewScanner(file)))
	file.Close()

	file, _ = os.Open("input.txt")
	fmt.Println("part2:", part2(bufio.NewScanner(file)))
	file.Close()
}
