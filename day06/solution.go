package main

import (
	"fmt"
	"log"
	"os"
)

func allCharsAreUnique(input []byte) bool {

	hasMatch := false

	for i := 0; i <= len(input) && !hasMatch; i++ {
		for j := i + 1; j < len(input) && !hasMatch; j++ {
			if j != i {
				hasMatch = input[i] == input[j]
			}
		}
	}

	return !hasMatch
}

func main() {
	filePath := os.Args[1]
	input, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Failed to read input file")
	}

	startOfPacket := 0

	for i := range input {
		if startOfPacket > 0 {
			continue
		}
		piece := input[i : i+4]
		if allCharsAreUnique(piece) {
			startOfPacket = i + 4
		}
	}

	fmt.Printf("Start of packet: %v\n", startOfPacket)

	startOfMessage := 0

	for i := range input {
		if startOfMessage > 0 {
			continue
		}
		piece := input[i : i+14]
		if allCharsAreUnique(piece) {
			startOfMessage = i + 14
		}
	}

	fmt.Printf("Start of message: %v\n", startOfMessage)
}
