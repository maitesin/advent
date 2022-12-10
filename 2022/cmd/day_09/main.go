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
	moves := strings.Split(strings.Trim(input, "\n"), "\n")
	head := Point{}
	tail := Point{}
	visited := map[Point]struct{}{
		tail: {},
	}

	for _, move := range moves {
		parts := strings.Split(move, " ")
		dir := parts[0]
		dis, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Panic(err)
		}
		head.Move(dir, dis)
		tail.Follow(head, visited)
	}

	return fmt.Sprint(len(visited))
}

func Task2(input string) string {
	moves := strings.Split(strings.Trim(input, "\n"), "\n")
	knots := []Point{
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {},
	}
	visited := map[Point]struct{}{
		Point{}: {},
	}

	for _, move := range moves {
		parts := strings.Split(move, " ")
		dir := parts[0]
		dis, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Panic(err)
		}
		knots[0].Move(dir, dis)
		for i := 1; i < len(knots)-1; i++ {
			knots[i].Follow(knots[i-1], map[Point]struct{}{})
		}
		knots[9].Follow(knots[8], visited)
	}

	return fmt.Sprint(len(visited))
}

type Point struct {
	X, Y int
}

func (p *Point) Move(dir string, dis int) {
	switch dir {
	case "U":
		p.Y += dis
	case "D":
		p.Y -= dis
	case "R":
		p.X += dis
	case "L":
		p.X -= dis
	default:
		log.Panic(fmt.Errorf("direction %q unrecognised", dir))
	}
}

func (p *Point) Follow(head Point, visited map[Point]struct{}) {
	disX := abs(p.X - head.X)
	disY := abs(p.Y - head.Y)

	if disY <= 1 && disX <= 1 {
		return
	}

	if disX > disY {
		// Lateral movement
		p.Y = head.Y
		if p.X > head.X {
			// Going Left
			for i := p.X - 1; i > head.X; i-- {
				visited[Point{X: i, Y: p.Y}] = struct{}{}
			}
			p.X = head.X + 1
		} else {
			// Going Right
			for i := p.X + 1; i < head.X; i++ {
				visited[Point{X: i, Y: p.Y}] = struct{}{}
			}
			p.X = head.X - 1
		}
	} else {
		// Vertical movement
		p.X = head.X
		if p.Y > head.Y {
			// Going Down
			for i := p.Y - 1; i > head.Y; i-- {
				visited[Point{X: p.X, Y: i}] = struct{}{}
			}
			p.Y = head.Y + 1
		} else {
			// Going up
			for i := p.Y + 1; i < head.Y; i++ {
				visited[Point{X: p.X, Y: i}] = struct{}{}
			}
			p.Y = head.Y - 1
		}
	}
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}
