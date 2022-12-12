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

	for i, row := range rows {
		m[i] = make([]Point, len(row))
		for j, height := range row {
			switch height {
			case 'S':
				S = Point{i, j, 'a'}
				m[i][j] = S
			case 'E':
				E = Point{i, j, 'z'}
				m[i][j] = E
			default:
				m[i][j] = Point{i, j, height}
			}
		}
	}

	paths := []Path{
		NewPath(S, E, map[Point]struct{}{}, 0),
	}

	for {
		sort.SliceStable(paths, func(i, j int) bool {
			return paths[i].Priority() < paths[j].Priority()
		})

		if len(paths) == 0 {
			log.Panic("You ran out of paths to explore")
		}
		p := paths[0]
		paths = paths[1:]

		if p.IsEnd() {
			return fmt.Sprint(p.Steps)
		}

		if p.Current.X > 0 {
			np := m[p.Current.X-1][p.Current.Y]
			if _, ok := p.Visited[np]; !ok && p.Current.StepSize(np) <= 1 {
				paths = append(paths, NewPath(np, E, p.Visited, p.Steps+1))
			}
		}
		if p.Current.Y > 0 {
			np := m[p.Current.X][p.Current.Y-1]
			if _, ok := p.Visited[np]; !ok && p.Current.StepSize(np) <= 1 {
				paths = append(paths, NewPath(np, E, p.Visited, p.Steps+1))
			}
		}
		if p.Current.X < len(m)-1 {
			np := m[p.Current.X+1][p.Current.Y]
			if _, ok := p.Visited[np]; !ok && p.Current.StepSize(np) <= 1 {
				paths = append(paths, NewPath(np, E, p.Visited, p.Steps+1))
			}
		}
		if p.Current.Y < len(m[0])-1 {
			np := m[p.Current.X][p.Current.Y+1]
			if _, ok := p.Visited[np]; !ok && p.Current.StepSize(np) <= 1 {
				paths = append(paths, NewPath(np, E, p.Visited, p.Steps+1))
			}
		}
	}

	log.Panic("You should not see this message")
	return ""
}

func Task2(input string) string {
	return input
}

type Point struct {
	X, Y  int
	Value rune
}

func (p Point) Distance(p2 Point) float64 {
	x := p.X - p2.X
	y := p.Y - p2.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (p Point) StepSize(p2 Point) rune {
	if p.Value < p2.Value {
		return p2.Value - p.Value
	}

	return p.Value - p2.Value
}

type Path struct {
	Steps   int
	End     Point
	Current Point
	Visited map[Point]struct{}
}

func NewPath(c, e Point, visited map[Point]struct{}, steps int) Path {
	p := Path{
		Steps:   steps,
		End:     e,
		Current: c,
	}

	p.Visited = make(map[Point]struct{}, len(visited)+1)
	p.Visited[c] = struct{}{}
	for key := range visited {
		p.Visited[key] = visited[key]
	}

	return p
}

func (p Path) Priority() float64 {
	return p.Current.Distance(p.End)
}

func (p Path) IsEnd() bool {
	return p.Current == p.End
}
