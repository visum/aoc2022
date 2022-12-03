package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//
// A, X | Rock     | 1
// B, Y | Paper    | 2
// C, Z | Scissors | 3
//
// Win 6
// Draw 3
// Lose 0

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	myMoveDecode := map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	theirMoveDecode := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	scoreMap := make(map[string]int)
	scoreMap["rock"] = 1
	scoreMap["paper"] = 2
	scoreMap["scissors"] = 3
	scoreMap["win"] = 6
	scoreMap["draw"] = 3
	scoreMap["lose"] = 0

	resultMap := map[string]map[string]string{
		"rock": map[string]string{
			"rock":     "draw",
			"paper":    "lose",
			"scissors": "win",
		},
		"paper": map[string]string{
			"rock":     "win",
			"paper":    "draw",
			"scissors": "lose",
		},
		"scissors": map[string]string{
			"rock":     "lose",
			"paper":    "win",
			"scissors": "draw",
		},
	}

	scoreTotal := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, " ")
		theirMove := theirMoveDecode[moves[0]]
		myMove := myMoveDecode[moves[1]]

		shapeScore := scoreMap[myMove]
		result := resultMap[myMove][theirMove]
		resultScore := scoreMap[result]

		fmt.Printf("%s %s %s \n", myMove, theirMove, result)

		scoreTotal = scoreTotal + shapeScore + resultScore
	}

	fmt.Println(scoreTotal)
}
