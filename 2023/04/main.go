package main

import (
	"embed"
	"fmt"
	"log"
	"math"
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

type card struct {
	Winner map[int]struct{}
	Values map[int]struct{}
}

func Task1(input string) string {
	cards := parseCards(strings.Split(input, "\n"))
	accum := 0

	for _, c := range cards {
		accum += calculateCardValue(c)
	}

	return fmt.Sprint(accum)
}

func parseCards(lines []string) []card {
	cards := make([]card, len(lines))

	for i, line := range lines {
		cards[i] = parseCard(line)
	}

	return cards
}

func parseCard(line string) card {
	sections := strings.Split(line, ":")
	parts := strings.Split(strings.TrimSpace(sections[1]), " | ")
	winners := readNumbers(parts[0])
	values := readNumbers(parts[1])
	return card{
		Winner: winners,
		Values: values,
	}
}

func readNumbers(raw string) map[int]struct{} {
	parts := strings.Split(raw, " ")
	numbers := map[int]struct{}{}

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			value, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			numbers[value] = struct{}{}
		}
	}

	return numbers
}

func calculateCardValue(c card) int {
	count := 0
	for k, _ := range c.Winner {
		if _, ok := c.Values[k]; ok {
			count++
		}
	}
	return int(math.Pow(2, float64(count-1)))
}

func Task2(input string) string {
	cards := parseCards(strings.Split(input, "\n"))
	cache := make(map[int]int, len(cards))
	needToProcess := make([]int, len(cards))

	for i, c := range cards {
		cache[i] = countWinningNumbers(c)
		needToProcess[i] = i
	}

	accum := 0

	for {
		if len(needToProcess) == 0 {
			break
		}
		pos := needToProcess[0]
		needToProcess = needToProcess[1:]
		value := cache[pos]
		for i := 1; i <= value; i++ {
			needToProcess = append(needToProcess, pos+i)
		}
		accum++
	}

	return fmt.Sprint(accum)
}

func countWinningNumbers(c card) int {
	count := 0
	for k, _ := range c.Winner {
		if _, ok := c.Values[k]; ok {
			count++
		}
	}
	return count
}
