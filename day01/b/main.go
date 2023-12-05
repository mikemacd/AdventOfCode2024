package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

var i = 0

type Datarows []Datarow

type Datarow interface{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		os.Exit(1)
	}
	data, _ := ReadInput(os.Args[1])

	rv, err := ProcessData(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("error:%+v\nresult:\n%+v\n", err, rv)

	os.Exit(0)
}

func ReadInput(filename string) (Datarows, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file:", filename)
		return nil, err
	}

	lines := bytes.Split(data, []byte("\n"))

	rv := make(Datarows, len(lines))

	idx := 0
	for i, line := range lines {
		if len(line) == 0 {
			idx++
			continue
		}

		datarow := transformInputLine(line)

		rv[i] = datarow

	}

	return rv, nil
}

func transformInputLine(line []byte) Datarow {
	var rv Datarow

	rv = string(line)

	return rv
}

func ProcessData(data Datarows) (interface{}, error) {

	var rv = int(0)

	for _, item := range data {
		rv += ProcessLine(item.(string))
	}

	return rv, nil
}

func ProcessLine(line string) int {
	r1 := regexp.MustCompile(`^.*?(one|two|three|four|five|six|seven|eight|nine|\d)`)
	r2 := regexp.MustCompile(`^.*?(enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|\d)`)

	r1m := r1.FindStringSubmatch(line)
	r2m := r2.FindStringSubmatch(reverse(line))

	ldi := decode(r1m[1])
	rdi := decode(r2m[1])

	rv := ldi*10 + rdi
	fmt.Printf("%d %s %s:%s %d:%d %d\n", i, line, r1m[1], r2m[1], ldi, rdi, rv)
	i = i + 1
	return rv
}

func decode(input string) int {
	switch input {
	case "1", "one", "eno":
		return 1
	case "2", "two", "owt":
		return 2
	case "3", "three", "eerht":
		return 3
	case "4", "four", "ruof":
		return 4
	case "5", "five", "evif":
		return 5
	case "6", "six", "xis":
		return 6
	case "7", "seven", "neves":
		return 7
	case "8", "eight", "thgie":
		return 8
	case "9", "nine", "enin":
		return 9
	}
	return 0
}

func reverse(input string) string {
	rv := []rune(input)
	for i, j := 0, len(rv)-1; i < j; i, j = i+1, j-1 {
		rv[i], rv[j] = rv[j], rv[i]
	}
	return string(rv)
}
