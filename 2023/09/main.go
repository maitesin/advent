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
	sequences := parseSequences(strings.Split(input, "\n"))
	accum := 0

	for _, sequence := range sequences {
		accum += findNextStep(sequence)
	}

	return fmt.Sprint(accum)
}

func parseSequences(lines []string) [][]int {
	sequences := make([][]int, len(lines))
	for i, line := range lines {
		sequences[i] = parseSequence(line)
	}
	return sequences
}

func parseSequence(line string) []int {
	values := strings.Split(strings.TrimSpace(line), " ")
	seq := make([]int, len(values))

	for i, raw := range values {
		value, err := strconv.Atoi(raw)
		if err != nil {
			panic(err)
		}
		seq[i] = value
	}

	return seq
}

func findNextStep(seq []int) int {
	steps := [][]int{seq}
	current := seq
	for {
		next := diffSequence(current)
		if hasAllZeros(next) {
			break
		}
		steps = append(steps, next)
		current = next
	}

	return calculateNextElement(steps)
}

func diffSequence(seq []int) []int {
	diff := make([]int, len(seq)-1)

	for i := 0; i < len(seq)-1; i++ {
		diff[i] = seq[i+1] - seq[i]
	}

	return diff
}

func hasAllZeros(seq []int) bool {
	for _, value := range seq {
		if value != 0 {
			return false
		}
	}
	return true
}

func calculateNextElement(steps [][]int) int {
	value := 0
	for i := len(steps) - 1; i >= 0; i-- {
		last := len(steps[i]) - 1
		value += steps[i][last]
	}
	return value
}

func Task2(input string) string {
	sequences := parseSequences(strings.Split(input, "\n"))
	accum := 0

	for _, sequence := range sequences {
		aux := findNextStep2(sequence)
		accum += aux
	}

	return fmt.Sprint(accum)
}

func findNextStep2(seq []int) int {
	steps := [][]int{seq}
	current := seq
	for {
		next := diffSequence(current)
		if hasAllZeros(next) {
			break
		}
		steps = append(steps, next)
		current = next
	}

	return calculateNextElement2(steps)
}

func calculateNextElement2(steps [][]int) int {
	value := 0
	for i := len(steps) - 1; i >= 0; i-- {
		value = steps[i][0] - value
	}
	return value
}
