package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elves := make([]int64, 0)

	var elf int64
	elf = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		// if an empty line
		if text == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}

		number, _ := strconv.ParseInt(text, 10, 64)
		elf = elf + number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] < elves[j]
	})

	// part one - top elf
	fmt.Println(elves[len(elves)-1])

	// part two - sum of top three elves
	var topThree int64
	topThree = 0

	for i := len(elves) - 3; i < len(elves); i++ {
		topThree = topThree + elves[i]
	}

	fmt.Println(topThree)
}
