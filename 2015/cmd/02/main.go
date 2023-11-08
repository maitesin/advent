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
	fmt.Printf("Task1: %q\n", Task1(strings.TrimSpace(string(input1))))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(strings.TrimSpace(string(input2))))
}

func Task1(input string) string {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		total += singleBox(line)
	}
	return fmt.Sprint(total)
}

func singleBox(input string) int {
	parts := strings.Split(input, "x")
	l, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	lw := l * w
	lh := l * h
	wh := w * h
	smallest := min(lw, lh, wh)
	return lw*2 + lh*2 + wh*2 + smallest
}

func Task2(input string) string {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		total += ribbon(line)
	}
	return fmt.Sprint(total)
}

func ribbon(input string) int {
	parts := strings.Split(input, "x")
	l, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	greatest := max(l, w, h)
	a, b := 0, 0
	if greatest == l {
		a, b = w, h
	} else if greatest == w {
		a, b = l, h
	} else {
		a, b = l, w
	}
	return a*2 + b*2 + l*w*h
}
