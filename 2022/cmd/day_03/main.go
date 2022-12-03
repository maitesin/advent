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
	rucksacks := strings.Split(strings.Trim(input, "\n"), "\n")
	prio := make([]rune, len(rucksacks))
	for i, rucksack := range rucksacks {
		mid := len(rucksack) / 2
		comp1 := []rune(rucksack[:mid+1])
		comp2 := []rune(rucksack[mid:])
		prio[i] = findShared(comp1, comp2)
	}
	return fmt.Sprint(calculatePriorities(prio))
}

func findShared(comp1, comp2 []rune) rune {
	for _, r1 := range comp1 {
		for _, r2 := range comp2 {
			if r1 == r2 {
				return r1
			}
		}
	}
	log.Panicf("You should not be here: %q %q", string(comp1), string(comp2))
	return ' '
}

func calculatePriorities(prio []rune) int {
	accum := 0
	for _, r := range prio {
		if r >= 'a' && r <= 'z' {
			accum += int(r - 'a' + 1)
		} else {
			accum += int(r - 'A' + 27)
		}
	}

	return accum
}

func Task2(input string) string {
	rucksacks := strings.Split(strings.Trim(input, "\n"), "\n")
	groups := make([][3]string, len(rucksacks)/3)
	for i, rucksack := range rucksacks {
		groups[i/3][i%3] = rucksack
	}
	prio := make([]rune, len(rucksacks)/3)
	for i, group := range groups {
		prio[i] = findUnique(group)
	}

	return fmt.Sprint(calculatePriorities(prio))
}

func findUnique(group [3]string) rune {
	for _, r1 := range group[0] {
		for _, r2 := range group[1] {
			if r1 == r2 {
				for _, r3 := range group[2] {
					if r1 == r3 {
						return r1
					}
				}
			}
		}
	}
	log.Panicf("You should not be here: %v", group)
	return ' '
}
