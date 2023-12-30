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

func Task1(input string) string {
	m := strings.Split(input, "\n")
	accum := 0
	m = expandMap(m)
	galaxies := findGalaxies(m)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			a := galaxies[j][0]
			b := galaxies[i][0]
			if a < b {
				a, b = b, a
			}
			accum += a - b
			a = galaxies[j][1]
			b = galaxies[i][1]
			if a < b {
				a, b = b, a
			}
			accum += a - b
		}
	}
	return fmt.Sprint(accum)
}

func expandMap(m []string) []string {
	rows := findEmptyRows(m)
	columns := findEmptyColumns(m)
	for i := 0; i < len(rows); i++ {
		m = append(append(m[:rows[i]], m[rows[i]]), m[rows[i]:]...)
	}
	for i := 0; i < len(columns); i++ {
		for j := 0; j < len(m); j++ {
			m[j] = m[j][:columns[i]] + string(m[j][columns[i]]) + m[j][columns[i]:]
		}
	}
	return m
}

func findEmptyRows(m []string) []int {
	var empty []int

	for i := len(m) - 1; i >= 0; i-- {
		if isRowEmpty(m[i]) {
			empty = append(empty, i)
		}
	}

	return empty
}

func isRowEmpty(row string) bool {
	for _, v := range row {
		if v != '.' {
			return false
		}
	}
	return true
}

func findEmptyColumns(m []string) []int {
	var empty []int

	for i := len(m) - 1; i >= 0; i-- {
		if isColumnEmpty(m, i) {
			empty = append(empty, i)
		}
	}

	return empty
}

func isColumnEmpty(m []string, i int) bool {
	for j := 0; j < len(m); j++ {
		if m[j][i] != '.' {
			return false
		}
	}
	return true
}

type pair [2]int

func findGalaxies(m []string) []pair {
	var coordinates []pair

	for i := range m {
		for j := range m[i] {
			if m[i][j] == '#' {
				coordinates = append(coordinates, [2]int{i, j})
			}
		}
	}

	return coordinates
}

func Task2(input string) string {
	return input
}
