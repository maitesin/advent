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
	fmt.Printf("Task1: %q\n", Task1(strings.TrimSpace(string(input1))))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(strings.TrimSpace(string(input2))))
}

type row struct {
	Source int
	Dest   int
	Range  int
}

type mm []row

func (m mm) Find(item int) int {
	for _, r := range m {
		if r.Source <= item && item <= r.Source+r.Range-1 {
			offset := item - r.Source
			return r.Dest + offset
		}
	}
	return item
}

func Task1(input string) string {
	blocks := strings.Split(input, "\n\n")
	seeds := readSeeds(blocks[0])

	planning := map[string]mm{}
	for i := 1; i < len(blocks); i++ {
		name, m := readBlock(blocks[i])
		planning[name] = m
	}

	seedLocations := make([]int, len(seeds))
	lookUpOrder := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for i, seed := range seeds {
		item := seed
		for _, lookUp := range lookUpOrder {
			step := planning[lookUp]
			item = step.Find(item)
		}
		seedLocations[i] = item
	}

	sort.Ints(seedLocations)

	return fmt.Sprint(seedLocations[0])
}

func readSeeds(line string) []int {
	parts := strings.Split(line[7:], " ")
	seeds := make([]int, len(parts))

	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		seeds[i] = value
	}

	return seeds
}

func readBlock(block string) (string, mm) {
	lines := strings.Split(block, "\n")
	name := strings.Split(lines[0], " ")
	m := make(mm, len(lines))

	for i := 1; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		dest, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		sour, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		rang, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		m = append(m, row{Source: sour, Dest: dest, Range: rang})
	}

	return name[0], m
}

func Task2(input string) string {
	return input
}
