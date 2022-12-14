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
	rocks := make([]RockFormation, len(lines))
	maxDepth := 0
	maxRight := 0
	for i, line := range lines {
		rocks[i] = NewRockFormation(line)
		if rocks[i].MaxDepth > maxDepth {
			maxDepth = rocks[i].MaxDepth
		}
		if rocks[i].MaxRight > maxRight {
			maxRight = rocks[i].MaxRight
		}
	}

	m := NewMap(maxDepth, maxRight)
	m.PlaceRocks(rocks...)

	accum := 0
	for m.DropSand() {
		accum++
	}

	return fmt.Sprint(accum)
}

func Task2(input string) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	rocks := make([]RockFormation, len(lines))
	maxDepth := 0
	maxRight := 0
	for i, line := range lines {
		rocks[i] = NewRockFormation(line)
		if rocks[i].MaxDepth > maxDepth {
			maxDepth = rocks[i].MaxDepth
		}
		if rocks[i].MaxRight > maxRight {
			maxRight = rocks[i].MaxRight
		}
	}

	m := NewMapWithFloor(maxDepth, maxRight)
	m.PlaceRocks(rocks...)

	accum := 0
	for m.DropSand() {
		accum++
	}

	return fmt.Sprint(accum)
}

type Rock struct {
	X, Y int
}

type RockFormation struct {
	Rocks    []Rock
	MaxDepth int
	MaxRight int
}

func NewRockFormation(line string) RockFormation {
	rocks := strings.Split(line, " -> ")
	rf := RockFormation{}
	maxDepth := 0
	maxRight := 0
	for i := range rocks {
		rock := strings.Split(rocks[i], ",")
		y, err := strconv.Atoi(rock[0])
		if err != nil {
			log.Panic(err)
		}
		x, err := strconv.Atoi(rock[1])
		if err != nil {
			log.Panic(err)
		}
		rf.Rocks = append(rf.Rocks, Rock{X: x, Y: y})
		if x > maxDepth {
			maxDepth = x
		}
		if y > maxRight {
			maxRight = y
		}
	}

	rf.MaxRight = maxRight
	rf.MaxDepth = maxDepth

	return rf
}

type Map struct {
	Tiles    [][]rune
	MaxDepth int
	MaxRight int
}

func NewMap(maxDepth, maxRight int) Map {
	m := Map{MaxDepth: maxDepth, MaxRight: maxRight}
	m.Tiles = make([][]rune, maxDepth+1)
	for i := range m.Tiles {
		m.Tiles[i] = make([]rune, maxRight+1)
	}
	return m
}

func NewMapWithFloor(maxDepth, maxRight int) Map {
	m := Map{MaxDepth: maxDepth + 2, MaxRight: maxRight + maxDepth}
	m.Tiles = make([][]rune, m.MaxDepth+1)
	for i := range m.Tiles {
		m.Tiles[i] = make([]rune, m.MaxRight+1)
	}

	for i := 0; i <= m.MaxRight; i++ {
		m.Tiles[len(m.Tiles)-1][i] = '#'
	}
	return m
}

func (m Map) PlaceRocks(formations ...RockFormation) {
	for _, formation := range formations {
		for i := 1; i < len(formation.Rocks); i++ {
			if formation.Rocks[i].Y == formation.Rocks[i-1].Y {
				if formation.Rocks[i].X < formation.Rocks[i-1].X {
					for j := formation.Rocks[i].X; j <= formation.Rocks[i-1].X; j++ {
						m.Tiles[j][formation.Rocks[i].Y] = '#'
					}
				} else {
					for j := formation.Rocks[i-1].X; j <= formation.Rocks[i].X; j++ {
						m.Tiles[j][formation.Rocks[i].Y] = '#'
					}
				}
			} else {
				if formation.Rocks[i].Y < formation.Rocks[i-1].Y {
					for j := formation.Rocks[i].Y; j <= formation.Rocks[i-1].Y; j++ {
						m.Tiles[formation.Rocks[i].X][j] = '#'
					}
				} else {
					for j := formation.Rocks[i-1].Y; j <= formation.Rocks[i].Y; j++ {
						m.Tiles[formation.Rocks[i].X][j] = '#'
					}
				}
			}
		}
	}
}

func (m Map) Display() {
	for i := range m.Tiles {
		for j := range m.Tiles[i] {
			fmt.Print(string(m.Tiles[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m Map) DropSand() bool {
	xPos := 500
	yPos := 0

	if m.Tiles[yPos][xPos] == 'o' {
		return false
	}

	for ; yPos < m.MaxDepth; yPos++ {
		if m.Tiles[yPos+1][xPos] == 'o' ||
			m.Tiles[yPos+1][xPos] == '#' {
			if m.Tiles[yPos+1][xPos-1] == 'o' ||
				m.Tiles[yPos+1][xPos-1] == '#' {
				if m.Tiles[yPos+1][xPos+1] == 'o' ||
					m.Tiles[yPos+1][xPos+1] == '#' {
					m.Tiles[yPos][xPos] = 'o'
					return true
				} else {
					xPos++
				}
			} else {
				xPos--
			}
		}
	}

	return false
}
