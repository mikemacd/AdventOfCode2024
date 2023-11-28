package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var totalScore int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// determine the score for this round
		var score int

		// read in the input
		// opponentMove, yourMove := scanner.Text()[0], scanner.Text()[1]
		opponentMove, yourMove := scanner.Text()[0], scanner.Text()[2]

		switch opponentMove {
		case 'A':
			switch yourMove {
			case 'X':
				score = 1+3
			case 'Y':
				score = 2+6
			case 'Z':
				score = 3
			}
		case 'B':
			switch yourMove {
			case 'X':
				score = 1
			case 'Y':
				score = 2+3
			case 'Z':
				score = 3+6
			}
		case 'C':
			switch yourMove {
			case 'X':
				score = 1+6
			case 'Y':
				score = 2
			case 'Z':
				score = 3+3
			}
		}

		fmt.Printf("YourMove: %c TheirMove: %c score:%v \n", yourMove, opponentMove, score)
		// add the score for this round to the total score
		totalScore += score
	}

	fmt.Println(totalScore)
}
