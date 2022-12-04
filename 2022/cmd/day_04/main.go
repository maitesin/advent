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
	pairs := strings.Split(strings.Trim(input, "\n"), "\n")
	ranges := make([][2]Range, len(pairs))

	for i, pair := range pairs {
		pairParts := strings.Split(pair, ",")
		ranges[i] = [2]Range{
			NewRange(pairParts[0]),
			NewRange(pairParts[1]),
		}
	}

	accum := 0
	for _, r := range ranges {
		if r[0].Overlaps(r[1]) || r[1].Overlaps(r[0]) {
			accum++
		}
	}

	return fmt.Sprint(accum)
}

func Task2(input string) string {
	pairs := strings.Split(strings.Trim(input, "\n"), "\n")
	ranges := make([][2]Range, len(pairs))

	for i, pair := range pairs {
		pairParts := strings.Split(pair, ",")
		ranges[i] = [2]Range{
			NewRange(pairParts[0]),
			NewRange(pairParts[1]),
		}
	}

	accum := 0
	for _, r := range ranges {
		if r[0].PartialOverlap(r[1]) || r[1].PartialOverlap(r[0]) {
			accum++
		}
	}

	return fmt.Sprint(accum)
}

type Range struct {
	Lower, Upper int
}

func NewRange(raw string) Range {
	pair := strings.Split(raw, "-")
	lower, err := strconv.Atoi(pair[0])
	if err != nil {
		log.Panic(err)
	}

	upper, err := strconv.Atoi(pair[1])
	if err != nil {
		log.Panic(err)
	}

	return Range{
		Lower: lower,
		Upper: upper,
	}
}

func (r Range) Contains(pos int) bool {
	return r.Lower <= pos && pos <= r.Upper
}

func (r Range) Overlaps(r2 Range) bool {
	return r.Contains(r2.Lower) && r.Contains(r2.Upper)
}

func (r Range) PartialOverlap(r2 Range) bool {
	return r.Contains(r2.Lower) || r.Contains(r2.Upper)
}
