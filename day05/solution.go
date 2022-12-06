package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type step struct {
	quantity    int
	source      int
	destination int
}

func parseInput(input string) ([]step, [][]string) {
	lines := strings.Split(input, "\n")

	stateLines := make([]string, 0)
	stepLines := make([]string, 0)

	readMode := "state"

	for _, v := range lines {
		if readMode == "state" {
			if v == "" {
				readMode = "steps"
				continue
			}
			stateLines = append(stateLines, v)
		}

		if readMode == "steps" {
			stepLines = append(stepLines, v)
		}

	}

	steps := make([]step, len(stepLines))

	for index, stepString := range stepLines {
		steps[index] = parseStep(stepString)
	}

	state := parseState(stateLines)

	return steps, state
}

func parseStep(input string) step {
	const quantityStart = 5
	quantityEnd := strings.Index(input, "from") - 1

	quantityString := input[quantityStart:quantityEnd]

	sourceStart := quantityEnd + 6
	sourceEnd := strings.Index(input, "to") - 1

	sourceString := input[sourceStart:sourceEnd]

	targetStart := sourceEnd + 4

	targetString := input[targetStart:]

	source, _ := strconv.ParseInt(sourceString, 10, 0)
	target, _ := strconv.ParseInt(targetString, 10, 0)
	quantity, _ := strconv.ParseInt(quantityString, 10, 0)

	return step{
		source:      int(source),
		destination: int(target),
		quantity:    int(quantity),
	}
}

func parseState(input []string) [][]string {
	lastLine := input[len(input)-1]
	numberOfStacks, _ := strconv.ParseInt(lastLine[len(lastLine)-2:len(lastLine)-1], 10, 0)
	stacks := make([][]string, numberOfStacks)
	for stackIndex := 0; stackIndex < int(numberOfStacks); stackIndex++ {
		stacks[stackIndex] = make([]string, 0)
	}

	for lineNumber := len(input) - 2; lineNumber >= 0; lineNumber-- {
		line := input[lineNumber]
		for stackIndex := 0; stackIndex < int(numberOfStacks); stackIndex++ {
			stackLabelPosition := (stackIndex * 4) + 1
			stackLabel := string(line[stackLabelPosition])
			if stackLabel != " " {
				stacks[stackIndex] = append(stacks[stackIndex], stackLabel)
			}
		}
	}

	return stacks
}

func operate(step step, state [][]string) {
	// fmt.Printf("\nStep %#v\n", step)

	sourceIndex := step.source - 1
	destinationIndex := step.destination - 1

	sourceStack := state[sourceIndex]
	// destinationStack := state[destinationIndex]

	divisionIndex := len(sourceStack) - step.quantity
	thingsToMove := sourceStack[divisionIndex:]
	// fmt.Printf("Moving %v from %v to %v\n", thingsToMove, sourceStack, destinationStack)

	for i := range thingsToMove {
		itemIndex := len(thingsToMove) - i - 1
		state[destinationIndex] = append(state[destinationIndex], thingsToMove[itemIndex])
	}

	state[sourceIndex] = sourceStack[:divisionIndex]

	// fmt.Printf("Result: %v\n", state)
}

func main() {
	filePath := os.Args[1]

	input, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to read input file")
	}

	steps, state := parseInput(string(input))

	fmt.Printf("starting state %v\n", state)

	for _, operation := range steps {
		operate(operation, state)
	}

	topThings := make([]string, 0)

	for _, v := range state {
		topThing := v[len(v)-1]
		topThings = append(topThings, topThing)
	}

	fmt.Printf("result: \n %v\n", state)
	fmt.Printf("Top things: %v\n", topThings)
}
