package main

import (
	"fmt"
	"os"
)

func solve(comms []uint8, marker_size int) int {
	for i := marker_size; i < len(comms); i++ {
		buffer := comms[i-marker_size : i]
		var packet_marker = make(map[any]bool)
		for _, c := range buffer {
			packet_marker[c] = true
		}
		if len(packet_marker) == marker_size {
			return i
		}
	}
	return 0
}

func main() {
	comms, _ := os.ReadFile("input.txt")

	fmt.Println("part1:", solve(comms, 4))
	fmt.Println("part2:", solve(comms, 14))
}

/* total := 0
buffer := ring.New(marker_size)

for index, char := range comms {
	total += 1
	buffer = buffer.Next()
	buffer.Value = char

	if index > marker_size-2 {
		var packet_marker = make(map[any]int)
		buffer.Do(func(c any) { packet_marker[c] += 1 })

		isMarker := true
		for _, c := range packet_marker {
			isMarker = isMarker && c < 2
		}

		if isMarker {
			break
		}
	}
}

return total */
