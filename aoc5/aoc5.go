package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func parse_stacks(file *os.File) ([9]*list.List, *bufio.Scanner) {
	var stacks [9]*list.List

	for index, _ := range stacks {
		stacks[index] = list.New()
	}

	scanner := bufio.NewScanner(file)
	for row := 0; row < 8; row++ {
		scanner.Scan()
		line := scanner.Text()
		for col := 0; col < 9; col++ {
			if c := string(line[1+col*4]); c != " " {
				stacks[col].PushFront(c)
			}
		}
	}

	scanner.Scan()
	scanner.Scan()

	return stacks, scanner
}

func part1(file *os.File) {
	stacks, scanner := parse_stacks(file)

	for scanner.Scan() {
		move := scanner.Text()
		var quantity, from, to int
		fmt.Sscanf(move, "move %d from %d to %d", &quantity, &from, &to)

		from_stack, to_stack := stacks[from-1], stacks[to-1]

		for i := 0; i < quantity; i++ {
			crate := from_stack.Back()

			to_stack.PushBack(crate.Value)
			from_stack.Remove(crate)
		}
	}

	fmt.Print("part1: ")
	for _, stack := range stacks {
		fmt.Print(stack.Back().Value)
	}
	fmt.Println("")
}

func part2(file *os.File) {
	stacks, scanner := parse_stacks(file)

	for scanner.Scan() {
		move := scanner.Text()
		var quantity, from, to int
		fmt.Sscanf(move, "move %d from %d to %d", &quantity, &from, &to)

		from_stack, to_stack, temp_list := stacks[from-1], stacks[to-1], list.New()

		for i := 0; i < quantity; i++ {
			crate := from_stack.Back()

			temp_list.PushFront(crate.Value)
			from_stack.Remove(crate)
		}

		to_stack.PushBackList(temp_list)
	}

	fmt.Print("part2: ")
	for _, stack := range stacks {
		fmt.Print(stack.Back().Value)
	}
	fmt.Println("")
}

func main() {
	file, _ := os.Open("input.txt")
	part1(file)
	file.Close()

	file, _ = os.Open("input.txt")
	part2(file)
	file.Close()
}
