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

func Task1(input string) string {
	lines := strings.Split(input, "\n")
	accum := 0
	for i := range lines {
		for j := range lines[i] {
			switch {
			case lines[i][j] == '.', isNumber(lines, i, j):
				continue
			default:
				accum += calculateAdjacentNumbers(lines, i, j)
			}
		}
	}
	return fmt.Sprint(accum)
}

func isNumber(lines []string, i, j int) bool {
	switch lines[i][j] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func calculateAdjacentNumbers(lines []string, i, j int) int {
	accum := 0
	endRow := len(lines[0]) - 1
	bottomRow := len(lines) - 1
	found := map[int]struct{}{}

	if i != 0 && j != 0 && isNumber(lines, i-1, j-1) {
		value := readNumber(lines, i-1, j-1)
		found[value] = struct{}{}
		accum += value
	}
	if j != 0 && isNumber(lines, i, j-1) {
		value := readNumber(lines, i, j-1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum += value
		}
	}
	if i != 0 && isNumber(lines, i-1, j) {
		value := readNumber(lines, i-1, j)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum += value
		}
	}
	if i != bottomRow && j != endRow && isNumber(lines, i+1, j+1) {
		value := readNumber(lines, i+1, j+1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum += value
		}
	}
	if i != bottomRow && isNumber(lines, i+1, j) {
		value := readNumber(lines, i+1, j)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum += value
		}
	}
	if j != endRow && isNumber(lines, i, j+1) {
		value := readNumber(lines, i, j+1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum += value
		}
	}
	if i != 0 && j != endRow && isNumber(lines, i-1, j+1) {
		value := readNumber(lines, i-1, j+1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum += value
		}
	}
	if i != bottomRow && j != 0 && isNumber(lines, i+1, j-1) {
		value := readNumber(lines, i+1, j-1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum += value
		}
	}
	return accum
}

func readNumber(lines []string, i, j int) int {
	line := lines[i]
	end := j
	for ; end < len(line); end++ {
		if !isNumber(lines, i, end) {
			end--
			break
		}
	}
	if end == len(line) {
		end--
	}
	start := j
	for ; 0 <= start; start-- {
		if !isNumber(lines, i, start) {
			break
		}
	}
	value, err := strconv.Atoi(line[start+1 : end+1])
	if err != nil {
		panic(err)
	}
	return value
}

func Task2(input string) string {
	lines := strings.Split(input, "\n")
	accum := 0
	for i := range lines {
		for j := range lines[i] {
			switch {
			case lines[i][j] == '*':
				accum += calculateAdjacentNumbers2(lines, i, j)
			}
		}
	}
	return fmt.Sprint(accum)
}

func calculateAdjacentNumbers2(lines []string, i, j int) int {
	accum := 1
	endRow := len(lines[0]) - 1
	bottomRow := len(lines) - 1
	found := map[int]struct{}{}

	if i != 0 && j != 0 && isNumber(lines, i-1, j-1) {
		value := readNumber(lines, i-1, j-1)
		found[value] = struct{}{}
		accum *= value
	}
	if j != 0 && isNumber(lines, i, j-1) {
		value := readNumber(lines, i, j-1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum *= value
		}
	}
	if i != 0 && isNumber(lines, i-1, j) {
		value := readNumber(lines, i-1, j)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum *= value
		}
	}
	if i != bottomRow && j != endRow && isNumber(lines, i+1, j+1) {
		value := readNumber(lines, i+1, j+1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum *= value
		}
	}
	if i != bottomRow && isNumber(lines, i+1, j) {
		value := readNumber(lines, i+1, j)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum *= value
		}
	}
	if j != endRow && isNumber(lines, i, j+1) {
		value := readNumber(lines, i, j+1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum *= value
		}
	}
	if i != 0 && j != endRow && isNumber(lines, i-1, j+1) {
		value := readNumber(lines, i-1, j+1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum *= value
		}
	}
	if i != bottomRow && j != 0 && isNumber(lines, i+1, j-1) {
		value := readNumber(lines, i+1, j-1)
		if _, ok := found[value]; !ok {
			found[value] = struct{}{}
			accum *= value
		}
	}
	if len(found) != 2 {
		return 0
	}
	return accum
}
