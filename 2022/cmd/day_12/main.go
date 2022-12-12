package main

import (
	"embed"
	"fmt"
	"log"
	"math"
	"sort"
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
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	S := Point{0, 0, 0}
	E := Point{0, 0, 0}
	m := make([][]Point, len(rows))
	stepsG := make(map[Point]int, len(rows)*len(rows[0]))

	for i, row := range rows {
		m[i] = make([]Point, len(row))
		for j, height := range row {
			switch height {
			case 'S':
				S = Point{i, j, 'a'}
				m[i][j] = S
				stepsG[S] = 0
			case 'E':
				E = Point{i, j, 'z'}
				m[i][j] = E
				stepsG[E] = math.MaxInt
			default:
				p := Point{i, j, height}
				m[i][j] = p
				stepsG[p] = math.MaxInt
			}
		}
	}

	return fmt.Sprint(dijkstra(stepsG, S, E, m, math.MaxInt))
}

func Task2(input string) string {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	S := Point{0, 0, 0}
	E := Point{0, 0, 0}
	m := make([][]Point, len(rows))
	stepsG := make(map[Point]int, len(rows)*len(rows[0]))
	var startingPoints []Point
	for i, row := range rows {
		m[i] = make([]Point, len(row))
		for j, height := range row {
			switch height {
			case 'a':
				p := Point{i, j, height}
				m[i][j] = p
				stepsG[p] = math.MaxInt
				startingPoints = append(startingPoints, p)
			case 'S':
				S = Point{i, j, 'a'}
				m[i][j] = S
				stepsG[S] = math.MaxInt
			case 'E':
				E = Point{i, j, 'z'}
				m[i][j] = E
				stepsG[E] = math.MaxInt
			default:
				p := Point{i, j, height}
				m[i][j] = p
				stepsG[p] = math.MaxInt
			}
		}
	}

	copySteps := make(map[Point]int, len(stepsG))
	for key := range stepsG {
		copySteps[key] = stepsG[key]
	}
	cutValue := dijkstra(copySteps, S, E, m, math.MaxInt)
	for i := range startingPoints {
		copySteps = make(map[Point]int, len(stepsG))
		for key := range stepsG {
			copySteps[key] = stepsG[key]
		}
		v := dijkstra(copySteps, startingPoints[i], E, m, cutValue)
		if v < cutValue {
			cutValue = v
		}
	}

	return fmt.Sprint(cutValue)
}

type Point struct {
	X, Y  int
	Value rune
}

func (p Point) StepSize(p2 Point) rune {
	return p2.Value - p.Value
}

func dijkstra(stepsG map[Point]int, S, E Point, m [][]Point, cutValue int) int {
	stepsG[S] = 0

	for {
		keys := make([]Point, 0, len(stepsG))

		for key := range stepsG {
			keys = append(keys, key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return stepsG[keys[i]] < stepsG[keys[j]]
		})

		if len(stepsG) == 0 {
			log.Panic("You ran out of paths to explore")
		}
		p := keys[0]
		if p == E {
			return stepsG[E]
		}
		if stepsG[p] >= cutValue {
			return math.MaxInt
		}

		if p.X > 0 {
			np := m[p.X-1][p.Y]
			if p.StepSize(np) <= 1 && stepsG[np] >= stepsG[p]+1 {
				stepsG[np] = stepsG[p] + 1
			}
		}
		if p.Y > 0 {
			np := m[p.X][p.Y-1]
			if p.StepSize(np) <= 1 && stepsG[np] >= stepsG[p]+1 {
				stepsG[np] = stepsG[p] + 1
			}
		}
		if p.X < len(m)-1 {
			np := m[p.X+1][p.Y]
			if p.StepSize(np) <= 1 && stepsG[np] >= stepsG[p]+1 {
				stepsG[np] = stepsG[p] + 1
			}
		}
		if p.Y < len(m[0])-1 {
			np := m[p.X][p.Y+1]
			if p.StepSize(np) <= 1 && stepsG[np] >= stepsG[p]+1 {
				stepsG[np] = stepsG[p] + 1
			}
		}
		delete(stepsG, keys[0])
	}
	return math.MaxInt
}
