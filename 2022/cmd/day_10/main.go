package main

import (
	"embed"
	"fmt"
	"github.com/maitesin/advent/2022/internal/cpu"
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
	fmt.Println("Task2:")
	fmt.Print(Task2(string(input2)))
}

func Task1(input string) string {
	accum := 0
	n := 0
	c := cpu.NewCPU()
	instructions := strings.Split(strings.Trim(input, "\n"), "\n")

	for i := 0; i <= 220; i++ {
		switch i {
		case 20, 60, 100, 140, 180, 220:
			accum += i * c.X()
		}
		if c.Tick() {
			ins, err := cpu.NewInstruction(instructions[n])
			if err != nil {
				log.Panic(err)
			}

			c.Compute(ins)
			n++
		}
	}

	return fmt.Sprint(accum)
}

func Task2(input string) string {
	screen := make([]string, 40*6)
	c := cpu.NewCPU()
	instructions := strings.Split(strings.Trim(input, "\n"), "\n")
	n := 0

	for i := 0; i < len(screen); i++ {
		if c.Tick() {
			ins, err := cpu.NewInstruction(instructions[n])
			if err != nil {
				log.Panic(err)
			}

			c.Compute(ins)
			n++
		}
		if c.X()-1 <= i%40 && i%40 <= c.X()+1 {
			screen[i] = "#"
		} else {
			screen[i] = "."
		}
	}

	output := ""
	for i := 0; i < 6; i++ {
		output += fmt.Sprintln(strings.Join(screen[i*40:(i+1)*40], ""))
	}
	return fmt.Sprint(output)
}
