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
	count := 0
	for _, line := range strings.Split(input, "\n") {
		line := strings.TrimSpace(line)
		if isNaughty1(line) {
			continue
		} else if isNice1(line) {
			count += 1
		}
	}
	return fmt.Sprint(count)
}

func isNaughty1(input string) bool {
	return strings.Contains(input, "ab") ||
		strings.Contains(input, "cd") ||
		strings.Contains(input, "pq") ||
		strings.Contains(input, "xy")
}

func isNice1(input string) bool {
	return atLeast3Vowels(input) &&
		twoEqualCharactersConsecutive(input)
}

func atLeast3Vowels(input string) bool {
	vowels := 0
	for _, c := range input {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		}
	}
	return vowels >= 3
}

func twoEqualCharactersConsecutive(input string) bool {
	first := rune(input[0])
	for _, c := range input[1:] {
		if first == c {
			return true
		}
		first = c
	}
	return false
}

func Task2(input string) string {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		line := strings.TrimSpace(line)
		if isNice2(line) {
			count += 1
		}
	}
	return fmt.Sprint(count)
}

func isNice2(input string) bool {
	return containsAPairOfTwoLettersNoOverlap(input) &&
		containsAtLeastOneLetterRepeatedWithAnotherLetterBetweenThem(input)
}

func containsAPairOfTwoLettersNoOverlap(input string) bool {
	if len(input) < 2 {
		return false
	}
	pos := 2
	pair := input[:pos]
	for ; pos+2 <= len(input); pos++ {
		if pair == input[pos:pos+2] {
			return true
		}
	}
	return containsAPairOfTwoLettersNoOverlap(input[1:])
}

func containsAtLeastOneLetterRepeatedWithAnotherLetterBetweenThem(input string) bool {
	if len(input) < 3 {
		return false
	}
	if input[0] == input[2] {
		return true
	}
	return containsAtLeastOneLetterRepeatedWithAnotherLetterBetweenThem(input[1:])
}
