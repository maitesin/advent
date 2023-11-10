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
	fmt.Printf("Task1: %q\n", Task1(strings.TrimSpace(string(input1))))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(strings.TrimSpace(string(input2))))
}

type grid [1000][1000]int

func (g *grid) performOperation1(oper operation) {
	for i := oper.Start.X; i <= oper.End.X; i++ {
		for j := oper.Start.Y; j <= oper.End.Y; j++ {
			switch oper.Type {
			case "on":
				g[i][j] = 1
			case "off":
				g[i][j] = 0
			default:
				if g[i][j] == 0 {
					g[i][j] = 1
				} else {
					g[i][j] = 0
				}
			}
		}
	}
}

func (g *grid) performOperation2(oper operation) {
	for i := oper.Start.X; i <= oper.End.X; i++ {
		for j := oper.Start.Y; j <= oper.End.Y; j++ {
			switch oper.Type {
			case "on":
				g[i][j] += 1
			case "off":
				g[i][j] -= 1
				if g[i][j] < 0 {
					g[i][j] = 0
				}
			default:
				g[i][j] += 2
			}
		}
	}
}

func (g *grid) countLightsOn() int {
	count := 0
	for i := range g {
		for j := range g[i] {
			count += g[i][j]
		}
	}
	return count
}

type point struct {
	X, Y int
}

type operation struct {
	Type  string
	Start point
	End   point
}

func Task1(input string) string {
	grid := grid{}
	operations := processOperations(input)
	for _, oper := range operations {
		grid.performOperation1(oper)
	}
	return fmt.Sprint(grid.countLightsOn())
}

func processOperations(input string) []operation {
	rawParts := strings.Split(input, "\n")
	var opers []operation

	for _, part := range rawParts {
		switch {
		case strings.HasPrefix(part, "turn on "):
			o, d := processPointsPart(part[8:])
			opers = append(opers, operation{Type: "on", Start: o, End: d})
		case strings.HasPrefix(part, "turn off "):
			o, d := processPointsPart(part[9:])
			opers = append(opers, operation{Type: "off", Start: o, End: d})
		case strings.HasPrefix(part, "toggle "):
			o, d := processPointsPart(part[7:])
			opers = append(opers, operation{Type: "toggle", Start: o, End: d})
		default:
			panic("You should not see this")
		}
	}

	return opers
}

func processPointsPart(input string) (point, point) {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " through ")
	return processPoint(parts[0]), processPoint(parts[1])
}

func processPoint(input string) point {
	parts := strings.Split(strings.TrimSpace(input), ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return point{
		X: x,
		Y: y,
	}
}

func Task2(input string) string {
	grid := grid{}
	operations := processOperations(input)
	for _, oper := range operations {
		grid.performOperation2(oper)
	}
	return fmt.Sprint(grid.countLightsOn())
}
