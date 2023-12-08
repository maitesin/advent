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
	times, distances := readTable(strings.Split(input, "\n"))
	total := 1

	for i := range times {
		total *= calculateRace(times[i], distances[i])
	}

	return fmt.Sprint(total)
}

func readTable(lines []string) ([]int, []int) {
	return readLine(lines[0][5:]), readLine(lines[1][9:])
}

func readLine(line string) []int {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	var values []int

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			value, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			values = append(values, value)
		}
	}

	return values
}

func calculateRace(t, d int) int {
	count := 0
	stop := false
	for i := 0; i < t; i++ {
		if i*(t-i) > d {
			stop = true
			count++
		} else {
			if stop {
				return count
			}
		}
	}
	return count
}

func Task2(input string) string {
	t, distance := readTable2(strings.Split(input, "\n"))
	return fmt.Sprint(calculateRace(t, distance))
}

func readTable2(lines []string) (int, int) {
	return readLine2(lines[0][5:]), readLine2(lines[1][9:])
}

func readLine2(line string) int {
	line = strings.ReplaceAll(line, " ", "")
	value, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return value
}
