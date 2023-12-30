package main

import (
	"embed"
	"fmt"
	"log"
	"math"
	"strings"
)

//go:embed inputs/*.txt
var inputsFS embed.FS

func main() {
	input1, err := inputsFS.ReadFile("inputs/1.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task1: %q\n", Task1(strings.TrimSpace(string(input1))))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(strings.TrimSpace(string(input2))))
}

type maze []string

func (m maze) findStart() (int, int) {
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 'S' {
				return i, j
			}
		}
	}
	panic("There should be an S somewhere in the maze")
}

func (m maze) FindLongestDistance() int {
	sI, sJ := m.findStart()
	visited := map[string]struct{}{}
	visited[fmt.Sprintf("%d,%d", sI, sJ)] = struct{}{}
	distance := 1
	lI, lJ, rI, rJ := m.findFirstMove(sI, sJ)

	visited[fmt.Sprintf("%d,%d", lI, lJ)] = struct{}{}
	visited[fmt.Sprintf("%d,%d", rI, rJ)] = struct{}{}
	for {
		lI, lJ = m.findNextMove(lI, lJ, visited)
		rI, rJ = m.findNextMove(rI, rJ, visited)
		distance++
		visited[fmt.Sprintf("%d,%d", lI, lJ)] = struct{}{}
		if _, ok := visited[fmt.Sprintf("%d,%d", rI, rJ)]; ok {
			break
		}
		visited[fmt.Sprintf("%d,%d", rI, rJ)] = struct{}{}
	}
	return distance
}

func (m maze) findFirstMove(i, j int) (int, int, int, int) {
	moves := make([]int, 0, 4)
	if i+1 < len(m) {
		switch m[i+1][j] {
		case '|', 'L', 'J':
			moves = append(moves, i+1, j)
		}
	}
	if i-1 >= 0 {
		switch m[i-1][j] {
		case '|', 'F', '7':
			moves = append(moves, i-1, j)
		}
	}
	if j+1 < len(m[i]) {
		switch m[i][j+1] {
		case '-', 'J', '7':
			moves = append(moves, i, j+1)
		}
	}
	if j-1 >= 0 {
		switch m[i][j-1] {
		case '-', 'F', 'L':
			moves = append(moves, i, j-1)
		}
	}
	if len(moves) != 4 {
		panic("Moves has more than 4 elements")
	}
	return moves[0], moves[1], moves[2], moves[3]
}

func (m maze) findNextMove(i, j int, visited map[string]struct{}) (int, int) {
	switch m[i][j] {
	case '|':
		if i+1 < len(m) {
			switch m[i+1][j] {
			case '|', 'L', 'J':
				if _, ok := visited[fmt.Sprintf("%d,%d", i+1, j)]; !ok {
					return i + 1, j
				}
			}
		}
		if i-1 >= 0 {
			switch m[i-1][j] {
			case '|', 'F', '7':
				if _, ok := visited[fmt.Sprintf("%d,%d", i-1, j)]; !ok {
					return i - 1, j
				}
			}
		}
	case '-':
		if j+1 < len(m[i]) {
			switch m[i][j+1] {
			case '-', 'J', '7':
				if _, ok := visited[fmt.Sprintf("%d,%d", i, j+1)]; !ok {
					return i, j + 1
				}
			}
		}
		if j-1 >= 0 {
			switch m[i][j-1] {
			case '-', 'F', 'L':
				if _, ok := visited[fmt.Sprintf("%d,%d", i, j-1)]; !ok {
					return i, j - 1
				}
			}
		}
	case 'L':
		if i-1 >= 0 {
			switch m[i-1][j] {
			case '|', 'F', '7':
				if _, ok := visited[fmt.Sprintf("%d,%d", i-1, j)]; !ok {
					return i - 1, j
				}
			}
		}
		if j+1 < len(m[i]) {
			switch m[i][j+1] {
			case '-', 'J', '7':
				if _, ok := visited[fmt.Sprintf("%d,%d", i, j+1)]; !ok {
					return i, j + 1
				}
			}
		}
	case 'J':
		if i-1 >= 0 {
			switch m[i-1][j] {
			case '|', 'F', '7':
				if _, ok := visited[fmt.Sprintf("%d,%d", i-1, j)]; !ok {
					return i - 1, j
				}
			}
		}
		if j-1 >= 0 {
			switch m[i][j-1] {
			case '-', 'F', 'L':
				if _, ok := visited[fmt.Sprintf("%d,%d", i, j-1)]; !ok {
					return i, j - 1
				}
			}
		}
	case '7':
		if i+1 < len(m) {
			switch m[i+1][j] {
			case '|', 'L', 'J':
				if _, ok := visited[fmt.Sprintf("%d,%d", i+1, j)]; !ok {
					return i + 1, j
				}
			}
		}
		if j-1 >= 0 {
			switch m[i][j-1] {
			case '-', 'F', 'L':
				if _, ok := visited[fmt.Sprintf("%d,%d", i, j-1)]; !ok {
					return i, j - 1
				}
			}
		}
	case 'F':
		if i+1 < len(m) {
			switch m[i+1][j] {
			case '|', 'L', 'J':
				if _, ok := visited[fmt.Sprintf("%d,%d", i+1, j)]; !ok {
					return i + 1, j
				}
			}
		}
		if j+1 < len(m[i]) {
			switch m[i][j+1] {
			case '-', 'J', '7':
				if _, ok := visited[fmt.Sprintf("%d,%d", i, j+1)]; !ok {
					return i, j + 1
				}
			}
		}
	}
	return -1, -1
}

func (m maze) FindMaze() [][2]int {
	sI, sJ := m.findStart()
	visited := map[string]struct{}{}
	visited[fmt.Sprintf("%d,%d", sI, sJ)] = struct{}{}
	lI, lJ, _, _ := m.findFirstMove(sI, sJ)
	visited[fmt.Sprintf("%d,%d", lI, lJ)] = struct{}{}
	var result [][2]int
	result = append(result, [2]int{sI, sJ})
	result = append(result, [2]int{lI, lJ})
	for {
		lI, lJ = m.findNextMove(lI, lJ, visited)
		visited[fmt.Sprintf("%d,%d", lI, lJ)] = struct{}{}
		if lI == -1 && lJ == -1 {
			break
		}
		result = append(result, [2]int{lI, lJ})
	}
	return result
}

func Task1(input string) string {
	m := parseMaze(input)
	return fmt.Sprint(m.FindLongestDistance())
}

func parseMaze(input string) maze {
	return strings.Split(input, "\n")
}

func Task2(input string) string {
	m := parseMaze(input)
	ma := m.FindMaze()
	leftist := findLestist(ma)
	top := findTop(ma, leftist)
	for ma[0][0] != top && ma[0][1] != leftist {
		ma = append(ma[1:], ma[0])
	}
	return fmt.Sprint(calculateShoelaceValue(ma))
}

func findLestist(m [][2]int) int {
	leftist := math.MaxInt
	for i := 0; i < len(m); i++ {
		if m[i][1] < leftist {
			leftist = m[i][1]
		}
	}
	return leftist
}

func findTop(m [][2]int, leftist int) int {
	top := math.MaxInt
	for i := 0; i < len(m); i++ {
		if m[i][1] == leftist && m[i][0] < top {
			top = m[i][0]
		}
	}
	return top
}

func calculateShoelaceValue(m [][2]int) int {
	accum := 0
	for i := 0; i < len(m)-1; i++ {
		accum += (m[i][0] * m[i+1][1]) - (m[i+1][0] * m[i][1])
	}
	return (accum + (m[0][0] * m[len(m)-1][1]) - (m[len(m)-1][0] * m[0][1])) / 2
}
