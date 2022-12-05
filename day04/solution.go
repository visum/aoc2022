package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func contains(start int64, end int64, test int64) bool {
	return test >= start && test <= end
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open((filePath))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	fullOverlaps := 0
	anyOverlaps := 0

	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		a1 := strings.Split(assignments[0], "-")
		a2 := strings.Split(assignments[1], "-")

		a1Start, _ := strconv.ParseInt(a1[0], 10, 0)
		a1End, _ := strconv.ParseInt(a1[1], 10, 0)
		a2Start, _ := strconv.ParseInt(a2[0], 10, 0)
		a2End, _ := strconv.ParseInt(a2[1], 10, 0)

		if (a2Start >= a1Start && a2End <= a1End) || (a1Start >= a2Start && a1End <= a2End) {
			fullOverlaps++
		}

		if contains(a1Start, a1End, a2Start) ||
			contains(a1Start, a1End, a2End) ||
			contains(a2Start, a2End, a1Start) ||
			contains(a2Start, a2End, a1End) {
			anyOverlaps++
		}

	}

	fmt.Printf("Complete overlaps: %v\nAny overlaps: %v\n", fullOverlaps, anyOverlaps)

}
