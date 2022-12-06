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
	for i := 3; i < len(input); i++ {
		if validateMarker(input[i-3:i+1]) {
			return fmt.Sprint(i + 1)
		}
	}

	log.Panic("We should have found some 4 character marker")
	return ""
}

func validateMarker(marker string) bool {
	found := map[rune]struct{}{}
	for i := range marker {
		if _, ok := found[rune(marker[i])]; ok {
			return false
		}

		found[rune(marker[i])] = struct{}{}
	}
	return true
}

func Task2(input string) string {
	for i := 13; i < len(input); i++ {
		if validateMarker(input[i-13:i+1]) {
			return fmt.Sprint(i + 1)
		}
	}

	log.Panic("We should have found some 14 character marker")
	return ""
}
