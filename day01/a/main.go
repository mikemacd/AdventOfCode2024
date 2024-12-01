package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

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

	left, right := []int{}, []int{}

	for _, item := range data {
		if item == nil {
			continue
		}
		is := string(item.(string))
		a, b := ProcessLine(is)
		left = append(left, a)
		right = append(right, b)
	}

	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	for i := range left {
		if left[i] > right[i] {
			rv += left[i] - right[i]
		}
		if left[i] < right[i] {
			rv += right[i] - left[i]
		}
	}
	return rv, nil
}

func ProcessLine(line string) (a, b int) {
	_, _ = fmt.Sscanf(line, "%d %d", &a, &b)
	return
}
