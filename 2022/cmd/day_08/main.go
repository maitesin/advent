package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs/*.txt
var inputsFS embed.FS

func main() {
	input1, err := inputsFS.ReadFile("inputs/1.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task1: %q\n", Task1(string(input1)))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(string(input2)))
}

func Task1(input string) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	trees := make([][]int, len(lines))

	for i, line := range lines {
		row := make([]int, len(line))
		for j, t := range line {
			v, err := strconv.Atoi(string(t))
			if err != nil {
				log.Panic(err)
			}
			row[j] = v
		}
		trees[i] = row
	}

	accum := 2*len(trees) + 2*len(trees[0]) - 4
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			if isVisible(trees, i, j) {
				accum += 1
			}
		}
	}

	return fmt.Sprint(accum)
}

func isVisible(trees [][]int, x, y int) bool {
	return top(trees, x, y) ||
		bottom(trees, x, y) ||
		left(trees, x, y) ||
		right(trees, x, y)
}

func top(trees [][]int, x, y int) bool {
	v := trees[x][y]
	for i := x - 1; i >= 0; i-- {
		if trees[i][y] >= v {
			return false
		}
	}

	return true
}

func bottom(trees [][]int, x, y int) bool {
	v := trees[x][y]
	for i := x + 1; i < len(trees); i++ {
		if trees[i][y] >= v {
			return false
		}
	}

	return true
}

func left(trees [][]int, x, y int) bool {
	v := trees[x][y]
	for i := y - 1; i >= 0; i-- {
		if trees[x][i] >= v {
			return false
		}
	}

	return true
}

func right(trees [][]int, x, y int) bool {
	v := trees[x][y]
	for i := y + 1; i < len(trees[x]); i++ {
		if trees[x][i] >= v {
			return false
		}
	}

	return true
}

func Task2(input string) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	trees := make([][]int, len(lines))

	for i, line := range lines {
		row := make([]int, len(line))
		for j, t := range line {
			v, err := strconv.Atoi(string(t))
			if err != nil {
				log.Panic(err)
			}
			row[j] = v
		}
		trees[i] = row
	}

	max := distance(trees, 0, 0)
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			v := distance(trees, i, j)
			if v > max {
				max = v
			}
		}
	}

	return fmt.Sprint(max)
}

func distance(trees [][]int, x, y int) int {
	return topDistance(trees, x, y) *
		bottomDistance(trees, x, y) *
		leftDistance(trees, x, y) *
		rightDistance(trees, x, y)
}

func topDistance(trees [][]int, x, y int) int {
	v := trees[x][y]
	d := 1
	for i := x - 1; i >= 0; i-- {
		if trees[i][y] >= v {
			return d
		}
		d++
	}

	return d - 1
}

func bottomDistance(trees [][]int, x, y int) int {
	v := trees[x][y]
	d := 1
	for i := x + 1; i < len(trees); i++ {
		if trees[i][y] >= v {
			return d
		}
		d++
	}

	return d - 1
}

func leftDistance(trees [][]int, x, y int) int {
	v := trees[x][y]
	d := 1
	for i := y - 1; i >= 0; i-- {
		if trees[x][i] >= v {
			return d
		}
		d++
	}

	return d - 1
}

func rightDistance(trees [][]int, x, y int) int {
	v := trees[x][y]
	d := 1
	for i := y + 1; i < len(trees[x]); i++ {
		if trees[x][i] >= v {
			return d
		}
		d++
	}

	return d - 1
}
