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

	badges := make([]rune, 0)
	group := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		group = append(group, scanner.Text())

		if len(group) == 3 {
			var badge rune = -1
			for _, r := range group[0] {
				if badge == -1 && strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
					badge = r
				}
			}
			badges = append(badges, badge)
			group = make([]string, 0)
		}

	}

	priorities := make([]rune, 0)
	for _, v := range badges {
		if v > 95 {
			priorities = append(priorities, v-96)
		} else {
			priorities = append(priorities, v-38)
		}
	}

	fmt.Printf("%v\n", badges)

	// does golang have a reduce?
	sum := 0
	for _, v := range priorities {
		sum += int(v)
	}

	fmt.Printf("%v\n", sum)
}
