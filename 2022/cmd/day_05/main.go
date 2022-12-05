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
	inputParts := strings.Split(strings.Trim(input, "\n"), "\n\n")
	stacks := readStacks(inputParts[0])
	stacks = rearrangeStacks(stacks, inputParts[1])
	tops := findTopStacks(stacks)
	return fmt.Sprint(strings.Join(tops, ""))
}

func readStacks(input string) [][]string {
	lines := strings.Split(input, "\n")

	stacks := make([][]string, len(lines[0])/3)
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 1; j < len(lines[i]); j += 4 {
			value := string(lines[i][j])
			if value != " " {
				stacks[(j-1)/3] = append(stacks[(j-1)/3], value)
			}
		}
	}

	var final [][]string
	for i := range stacks {
		if len(stacks[i]) != 0 {
			final = append(final, stacks[i])
		}
	}

	return final
}

func rearrangeStacks(stacks [][]string, input string) [][]string {
	moves := strings.Split(input, "\n")

	for _, move := range moves {
		moveParts := strings.Split(move, " ")
		quantity, err := strconv.Atoi(moveParts[1])
		if err != nil {
			log.Panic(err)
		}
		from, err := strconv.Atoi(moveParts[3])
		if err != nil {
			log.Panic(err)
		}
		to, err := strconv.Atoi(moveParts[5])
		if err != nil {
			log.Panic(err)
		}
		for i := 0; i < quantity; i++ {
			stacks[from-1], stacks[to-1] = moveBetweenStacks(stacks[from-1], stacks[to-1])
		}
	}

	return stacks
}

func moveBetweenStacks(src []string, dst []string) ([]string, []string) {
	dst = append(dst, src[len(src)-1])
	src = src[:len(src)-1]
	return src, dst
}

func findTopStacks(stacks [][]string) []string {
	tops := make([]string, len(stacks))

	for i, stack := range stacks {
		tops[i] = stack[len(stack)-1]
	}

	return tops
}

func Task2(input string) string {
	inputParts := strings.Split(strings.Trim(input, "\n"), "\n\n")
	stacks := readStacks(inputParts[0])
	stacks = rearrangeStacks9001(stacks, inputParts[1])
	tops := findTopStacks(stacks)
	return fmt.Sprint(strings.Join(tops, ""))
}

func rearrangeStacks9001(stacks [][]string, input string) [][]string {
	moves := strings.Split(input, "\n")

	for _, move := range moves {
		moveParts := strings.Split(move, " ")
		quantity, err := strconv.Atoi(moveParts[1])
		if err != nil {
			log.Panic(err)
		}
		from, err := strconv.Atoi(moveParts[3])
		if err != nil {
			log.Panic(err)
		}
		to, err := strconv.Atoi(moveParts[5])
		if err != nil {
			log.Panic(err)
		}
		stacks[from-1], stacks[to-1] = moveBetweenStacks9001(stacks[from-1], stacks[to-1], quantity)
	}

	return stacks
}

func moveBetweenStacks9001(src []string, dst []string, quantity int) ([]string, []string) {
	dst = append(dst, src[len(src)-quantity:]...)
	src = src[:len(src)-quantity]
	return src, dst
}
