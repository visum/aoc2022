package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//
// A | Rock     | 1
// B | Paper    | 2
// C | Scissors | 3
//
// Z Win 6
// Y Draw 3
// X Lose 0

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	theirMoveDecode := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	myMoveMap := map[string]map[string]string{
		"rock": map[string]string{
			"X": "scissors",
			"Y": "rock",
			"Z": "paper",
		},
		"paper": map[string]string{
			"X": "rock",
			"Y": "paper",
			"Z": "scissors",
		},
		"scissors": map[string]string{
			"X": "paper",
			"Y": "scissors",
			"Z": "rock",
		},
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
		parts := strings.Split(line, " ")
		theirMove := theirMoveDecode[parts[0]]
		expectedResult := parts[1]

		myMove := myMoveMap[theirMove][expectedResult]

		shapeScore := scoreMap[myMove]
		result := resultMap[myMove][theirMove]
		resultScore := scoreMap[result]

		fmt.Printf("%s %s %s \n", myMove, theirMove, result)

		scoreTotal = scoreTotal + shapeScore + resultScore
	}

	fmt.Println(scoreTotal)
}
