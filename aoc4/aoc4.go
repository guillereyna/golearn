package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(file *os.File) int {
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")
		f, s := strings.Split(elves[0], "-"), strings.Split(elves[1], "-")

		fstart, _ := strconv.ParseInt(f[0], 10, 64)
		fend, _ := strconv.ParseInt(f[1], 10, 64)
		sstart, _ := strconv.ParseInt(s[0], 10, 64)
		send, _ := strconv.ParseInt(s[1], 10, 64)

		var first_sections, second_sections, intersections [100]bool

		for i := fstart; i < fend+1; i++ {
			first_sections[i] = true
		}

		for i := sstart; i < send+1; i++ {
			second_sections[i] = true
		}

		for index, _ := range intersections {
			intersections[index] = first_sections[index] && second_sections[index]
		}

		res := intersections == first_sections || intersections == second_sections

		if res {
			total += 1
		}
	}
	return total
}

func part2(file *os.File) int {
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")
		f, s := strings.Split(elves[0], "-"), strings.Split(elves[1], "-")

		fstart, _ := strconv.ParseInt(f[0], 10, 64)
		fend, _ := strconv.ParseInt(f[1], 10, 64)
		sstart, _ := strconv.ParseInt(s[0], 10, 64)
		send, _ := strconv.ParseInt(s[1], 10, 64)

		var first_sections, second_sections [100]bool

		for i := fstart; i < fend+1; i++ {
			first_sections[i] = true
		}

		for i := sstart; i < send+1; i++ {
			second_sections[i] = true
		}

		res := false
		for index, _ := range first_sections {
			res = res || first_sections[index] && second_sections[index]
		}

		if res {
			total += 1
		}
	}
	return total
}

func main() {
	file, _ := os.Open("input.txt")
	fmt.Println("part1:", part1(file))
	file.Close()

	file, _ = os.Open("input.txt")
	fmt.Println("part2:", part2(file))
	file.Close()
}
