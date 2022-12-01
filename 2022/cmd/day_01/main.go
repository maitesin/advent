package main

import (
	"embed"
	"fmt"
	"log"
	"sort"
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
	elvesStrings := strings.Split(input, "\n\n")
	elvesValues := make([]int, len(elvesStrings))

	for i, elfString := range elvesStrings {
		elfValues := 0
		for _, valueString := range strings.Split(elfString, "\n") {
			value, err := strconv.Atoi(valueString)
			if err != nil {
				log.Panic(err)
			}
			elfValues += value
		}
		elvesValues[i] = elfValues
	}

	sort.Ints(elvesValues)
	return fmt.Sprint(elvesValues[len(elvesValues)-1])
}

func Task2(input string) string {
	elvesStrings := strings.Split(input, "\n\n")
	elvesValues := make([]int, len(elvesStrings))

	for i, elfString := range elvesStrings {
		elfValues := 0
		for _, valueString := range strings.Split(elfString, "\n") {
			value, err := strconv.Atoi(valueString)
			if err != nil {
				log.Panic(err)
			}
			elfValues += value
		}
		elvesValues[i] = elfValues
	}

	sort.Ints(elvesValues)
	l := len(elvesValues)

	return fmt.Sprint(elvesValues[l-1] + elvesValues[l-2] + elvesValues[l-3])
}
