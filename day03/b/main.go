package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Keep track of the sum of priorities
	sum := 0

	// Read the input line by line
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		// Split the line into the two compartments
		half := len(line) / 2
		firstCompartment := line[:half]
		secondCompartment := line[half:]

		if i%3 == 0 {
			// First rucksack in group
			// Find the badge for this group
			badge := rune((i/3) + 'a')
			if !strings.ContainsRune(firstCompartment, badge) || !strings.ContainsRune(secondCompartment, badge) {
				sum += getPriority(badge)
			}
		} else {
			// Second or third rucksack in group
			// Find the item that appears in both compartments
			for _, c := range firstCompartment {
				if strings.ContainsRune(secondCompartment, c) {
					sum += getPriority(c)
					break
				}
			}
		}
	}

	fmt.Println(sum)
}

func getPriority(c rune) int {
	// Lowercase item types a through z have priorities 1 through 26
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}

	// Uppercase item types A through Z have priorities 27 through 52
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	}

	return 0
}