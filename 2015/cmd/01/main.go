package main

import (
	"embed"
	"fmt"
	"log"
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
	counter := 0
	for _, c := range input {
		if c == '(' {
			counter += 1
		} else if c == ')' {
			counter -= 1
		} else {
			panic(c)
		}
	}
	return fmt.Sprint(counter)
}

func Task2(input string) string {
	counter := 0
	for pos, c := range input {
		if c == '(' {
			counter += 1
		} else if c == ')' {
			counter -= 1
		} else {
			panic(c)
		}
		if counter < 0 {
			return fmt.Sprint(pos + 1)
		}
	}
	return fmt.Sprint(len(input))
}
