package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type elves []singleElf
type singleElf struct {
	items []int
}

func main() {
	elves := readInput()

	idx, cal := findBiggest(elves)

	fmt.Printf("%d has the most calories %d \n", idx, cal)
	os.Exit(0)
}

func readInput() elves {
	e := elves{}

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return elves{}
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	lines := bytes.Split(data, []byte("\n"))
	idx := 0
	for i, line := range lines {
		if len(line) == 0 {
			idx++
			continue
		}
		num, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatalf("Can't parse number on line %d: %v\n", i, line)
		}
		if len(e)-1 < idx {
			e = append(e, singleElf{})
		}
		e[idx].items = append(e[idx].items, num)
	}

	return e
}

func findBiggest(e elves) (int, int) {
	biggest := -1
	biggestCal := -1

	for i, v := range e {
		sum := 0
		for _, c := range v.items {
			sum += c
		}
		if sum > biggestCal {
			biggestCal = sum
			biggest = i
		}
	}
	return biggest, biggestCal
}
