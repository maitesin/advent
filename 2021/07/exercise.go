package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateMinFuelConsume(positions []int, distance int) int {
	maxPos := 0
	for _, position := range positions {
		if position >= maxPos {
			maxPos = position
		}
	}

	cache := make(map[int]int, maxPos)
	accum := 0
	for i :=0; i<= maxPos; i++ {
		if distance == 0 {
			cache[i] = i
		} else {
			accum += i
			cache[i] = accum
		}
	}

	consume := math.MaxInt
	for i := 0; i <= maxPos; i++ {
		newConsume := calculateDifferences(positions, i, cache)
		if newConsume < consume {
			consume = newConsume
		}
	}

	return consume
}

func calculateDifferences(positions []int, value int, distances map[int]int) int {
	accum := 0
	for _, position := range positions {
		tmp := value - position
		if tmp < 0 {
			tmp = position - value
		}
		accum += distances[tmp]
	}

	return accum
}

func main() {
	f1 := readInput()
	positions := initialPositions(f1)

	fmt.Printf("Answer #1: %d\n", calculateMinFuelConsume(positions, 0))
	fmt.Printf("Answer #2: %d\n", calculateMinFuelConsume(positions, 1))
}

func initialPositions(closer io.ReadCloser) []int {
	buffer := bufio.NewReader(closer)

	var positions []int

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
		positions = append(positions, stateInt)
	}

	return positions
}

func readInput() io.ReadCloser {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	return file
}
