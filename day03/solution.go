package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	duplicates := make([]rune, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		half := len(line) / 2

		firstCompartment := line[0:half]
		secondCompartment := line[half:]
		var duplicate rune

		for _, v := range firstCompartment {
			if strings.ContainsRune(secondCompartment, v) {
				duplicate = v
			}
		}

		duplicates = append(duplicates, duplicate)
	}

	priorities := make([]rune, 0)
	for _, v := range duplicates {
		if v > 95 {
			priorities = append(priorities, v-96)
		} else {
			priorities = append(priorities, v-38)
		}
	}

	// does golang have a reduce?
	sum := 0
	for _, v := range priorities {
		sum += int(v)
	}

	fmt.Printf("%v\n", sum)
}
