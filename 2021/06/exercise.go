package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type cacheIndex struct {
	state int
	days  int
}

func iterateStatesForDays(states []int, days int) int {
	accum := 0
	cache := make(map[cacheIndex]int)

	for _, state := range states {
		accum += calculateDescendants(state, days, cache)
	}

	return accum
}

func calculateDescendants(state, days int, cache map[cacheIndex]int) int {
	remainingDays := days - state - 1
	if total, ok := cache[cacheIndex{state: state, days: remainingDays}]; ok {
		return total
	}

	childrenTotal := 1
	for i := remainingDays; i >= 0; i -= 7 {
		childrenTotal += calculateDescendants(8, i, cache)
	}

	cache[cacheIndex{state: state, days: remainingDays}] = childrenTotal

	return childrenTotal
}

func main() {
	f1 := readInput()
	states := initialState(f1)

	fmt.Printf("Answer #1: %d\n", iterateStatesForDays(states, 80))
	fmt.Printf("Answer #2: %d\n", iterateStatesForDays(states, 256))
}

func initialState(closer io.ReadCloser) []int {
	buffer := bufio.NewReader(closer)

	var initialState []int

	line, _, err := buffer.ReadLine()
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}
	states := strings.Split(string(line), ",")
	for _, state := range states {
		stateInt, err := strconv.Atoi(state)
		if err != nil {
			panic(fmt.Sprintf("%+v\n", err))
		}
		initialState = append(initialState, stateInt)
	}

	return initialState
}

func readInput() io.ReadCloser {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	return file
}
