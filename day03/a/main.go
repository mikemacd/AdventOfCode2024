package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		half := len(line) / 2
		firstCompartment := line[:half]
		secondCompartment := line[half:]
		for _, c := range firstCompartment {
			if strings.ContainsRune(secondCompartment, c) {
				sum += getPriority(c)
				break
			}
		}
	}
	fmt.Println(sum)
}

func getPriority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	}
	return 0
}