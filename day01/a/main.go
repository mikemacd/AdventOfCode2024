package main

import (
	"bytes"
	"fmt"
	"os"
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
fmt.Printf("LinesLen:%d\n",len(lines))
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

	return rv
}

func ProcessData(data Datarows) (interface{}, error) {
	var rv interface{}

	return rv, nil
}
