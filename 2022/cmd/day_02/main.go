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
	fmt.Printf("Task1: %q\n", Task1(string(input1)))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(string(input2)))
}

func Task1(input string) string {
	values := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	decOp := map[string]string{
		"A": "R",
		"B": "P",
		"C": "S",
	}

	decU := map[string]string{
		"X": "R",
		"Y": "P",
		"Z": "S",
	}

	score := 0
	plays := strings.Split(strings.Trim(input, "\n"), "\n")
	for _, play := range plays {
		sections := strings.Split(play, " ")
		op, u := sections[0], sections[1]
		outcome := calculateOutcomeTask1(decOp[op], decU[u])
		score += outcome + values[u]
	}

	return fmt.Sprint(score)
}

func calculateOutcomeTask1(op, u string) int {
	switch {
	case op == u:
		return 3
	case op == "R" && u == "P",
		op == "P" && u == "S",
		op == "S" && u == "R":
		return 6
	default:
		return 0
	}
}

func Task2(input string) string {
	values := map[string]int{
		"R": 1,
		"P": 2,
		"S": 3,
	}

	decOp := map[string]string{
		"A": "R",
		"B": "P",
		"C": "S",
	}

	decU := map[string]string{
		"X": "L",
		"Y": "D",
		"Z": "W",
	}

	decUValue := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	score := 0
	plays := strings.Split(strings.Trim(input, "\n"), "\n")
	for _, play := range plays {
		sections := strings.Split(play, " ")
		op, u := sections[0], sections[1]
		outcome := calculateOutcomeTask2(decOp[op], decU[u])
		score += decUValue[u] + values[outcome]
	}

	return fmt.Sprint(score)
}

func calculateOutcomeTask2(op, u string) string {
	if u == "L" {
		switch op {
		case "R":
			return "S"
		case "P":
			return "R"
		default:
			return "P"
		}
	} else if u == "D" {
		switch op {
		case "R":
			return "R"
		case "P":
			return "P"
		default:
			return "S"
		}
	} else { // u == "W"
		switch op {
		case "R":
			return "P"
		case "P":
			return "S"
		default:
			return "R"
		}
	}
}
