package main

import (
	"embed"
	"fmt"
	"log"
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

type Position struct {
	X, Y int
}

func Task1(input string) string {
	pos := Position{}
	visited := make(map[Position]struct{})
	visited[pos] = struct{}{}
	for _, move := range input {
		switch move {
		case '^':
			pos.Y += 1
		case 'v':
			pos.Y -= 1
		case '>':
			pos.X += 1
		case '<':
			pos.X -= 1
		default:
			panic("You should not be here")
		}
		visited[pos] = struct{}{}
	}
	return fmt.Sprint(len(visited))
}

func Task2(input string) string {
	santa := Position{}
	robot := Position{}
	visited := make(map[Position]struct{})
	visited[santa] = struct{}{}
	for i, move := range input {
		pos := &robot
		if i%2 == 0 {
			pos = &santa
		}
		switch move {
		case '^':
			pos.Y += 1
		case 'v':
			pos.Y -= 1
		case '>':
			pos.X += 1
		case '<':
			pos.X -= 1
		default:
			panic("You should not be here")
		}
		visited[*pos] = struct{}{}
	}
	return fmt.Sprint(len(visited))
}
