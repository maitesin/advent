package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f1 := readInput()
	values := listOfIntegers(f1)

	prev := values[0]
	counter := 0
	for _, value := range values[1:] {
		if value > prev {
			counter++
		}
		prev = value
	}
	fmt.Printf("Answer #1: %d\n", counter)

	beg := values[0]
	mid := values[1]
	end := values[2]
	counter = 0
	for _, value := range values[3:] {
		prev := beg + mid + end
		curr := mid + end + value
		if curr > prev {
			counter++
		}
		beg = mid
		mid = end
		end = value
	}

	fmt.Printf("Answer #2: %d\n", counter)
}

func listOfIntegers(closer io.ReadCloser) []int {
	buffer := bufio.NewReader(closer)

	var values []int
	for {
		valueStr, _, err := buffer.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("%+v\n", err))
		}
		valueInt, err := strconv.Atoi(strings.TrimSpace(string(valueStr)))
		if err != nil {
			panic(fmt.Sprintf("%+v\n", err))
		}
		values = append(values, valueInt)
	}
	return values
}


func readInput() io.ReadCloser {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	return file
}