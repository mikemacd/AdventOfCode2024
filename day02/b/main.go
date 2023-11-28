package main

import (
	"bufio"
	"fmt"
	"os"
)

type move int
type score int

const (
	rock     = 1
	paper    = 2
	scissors = 3

	loss = 0 
	draw = 3 
	win  = 6 
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
		case 'A': // rock
			switch yourMove {
			case 'X': // lose
				score = int(scissors) + int(loss)
			case 'Y': // draw
				score = int(rock) + int(draw)
			case 'Z': // win
				score = int(paper) + int(win)
			}
		case 'B': // paper
			switch yourMove {
			case 'X': // lose
				score = int(rock) + int(loss)
			case 'Y': // draw
				score = int(paper) + int(draw)
			case 'Z': // win
				score = int(scissors) + int(win)
			}
		case 'C': // scissors
			switch yourMove {
			case 'X': // lose
				score = int(paper) + int(loss)
			case 'Y': // draw
				score = int(scissors) + int(draw)
			case 'Z': // win
				score = int(rock) + int(win)
			}
		}

		fmt.Printf("YourMove: %c TheirMove: %c score:%v \n", yourMove, opponentMove, score)
		// add the score for this round to the total score
		totalScore += score
	}

	fmt.Println(totalScore)
}
