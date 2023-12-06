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

type color string

const red color = "red"
const blue color = "blue"
const green color = "green"

type game struct {
	Number int
	Plays  []map[color]int
}

func Task1(input string) string {
	lines := strings.Split(input, "\n")
	games := parseGames(lines)
	accum := 0
	redValue := 12
	blueValue := 14
	greenValue := 13

	for _, g := range games {
		if isGamePossible(g, redValue, blueValue, greenValue) {
			accum += g.Number
		}
	}

	return fmt.Sprint(accum)
}

func parseGames(lines []string) []game {
	games := make([]game, len(lines))
	for i, line := range lines {
		games[i] = parseGame(line)
	}
	return games
}

func parseGame(line string) game {
	sections := strings.Split(line, ": ")
	number, err := strconv.Atoi(sections[0][5:])
	if err != nil {
		panic(err)
	}

	playParts := strings.Split(sections[1], "; ")
	plays := make([]map[color]int, len(playParts))
	for i, parts := range playParts {
		plays[i] = parsePlay(strings.TrimSpace(parts))
	}

	return game{
		Number: number,
		Plays:  plays,
	}
}

func parsePlay(play string) map[color]int {
	playParts := strings.Split(play, ", ")
	m := make(map[color]int, 3)

	for _, playPart := range playParts {
		parts := strings.Split(playPart, " ")
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		switch parts[1] {
		case "red":
			m[red] = value
		case "blue":
			m[blue] = value
		case "green":
			m[green] = value
		default:
			panic(fmt.Sprintf("Unknown color %q", parts[1]))
		}
	}

	return m
}

func isGamePossible(p game, r, b, g int) bool {
	for _, play := range p.Plays {
		if !isPlayPossible(play, r, b, g) {
			return false
		}
	}
	return true
}

func isPlayPossible(p map[color]int, r, b, g int) bool {
	if v, ok := p[red]; ok {
		if v > r {
			return false
		}
	}
	if v, ok := p[blue]; ok {
		if v > b {
			return false
		}
	}
	if v, ok := p[green]; ok {
		if v > g {
			return false
		}
	}
	return true
}

func Task2(input string) string {
	lines := strings.Split(input, "\n")
	games := parseGames(lines)
	accum := 0

	for _, g := range games {
		accum += calculateGamePower(g)
	}

	return fmt.Sprint(accum)
}

func calculateGamePower(g game) int {
	redMin := findMinColor(g.Plays, red)
	greenMin := findMinColor(g.Plays, green)
	blueMin := findMinColor(g.Plays, blue)
	return redMin * greenMin * blueMin
}

func findMinColor(plays []map[color]int, c color) int {
	min := 0

	for _, play := range plays {
		if v, ok := play[c]; ok {
			if v > min {
				min = v
			}
		}
	}

	return min
}
