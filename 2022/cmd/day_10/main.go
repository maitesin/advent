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
	instructions := strings.Split(strings.Trim(input, "\n"), "\n")
	cycle := 1
	x := 1
	prevX := x

	stops := []int{20, 60, 100, 140, 180, 220}
	accum := 0
	sI := 0

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		oper := parts[0]
		value := 0
		var err error
		if len(parts) > 1 {
			value, err = strconv.Atoi(parts[1])
			if err != nil {
				log.Panic(err)
			}
		}

		switch oper {
		case "addx":
			cycle += 2
			x += value
		case "noop":
			cycle++
		default:
			log.Panic("You should not see this")
		}

		if sI < len(stops) {
			switch {
			case cycle == stops[sI]:
				accum += x * stops[sI]
				sI++
			case cycle > stops[sI]:
				accum += prevX * stops[sI]
				sI++
			}
		} else {
			break
		}

		prevX = x
	}

	return fmt.Sprint(accum)
}

func Task2(input string) string {
	instructions := strings.Split(strings.Trim(input, "\n"), "\n")
	n := 0
	cycle := 1
	x := 1
	prevX := x
	wait := 0

	screen := make([]string, 40*6)

	for {
		if wait == 0 {
			prevX = x
			if n >= len(instructions) {
				break
			}
			parts := strings.Split(instructions[n], " ")
			oper := parts[0]
			value := 0
			var err error
			if len(parts) > 1 {
				value, err = strconv.Atoi(parts[1])
				if err != nil {
					log.Panic(err)
				}
			}

			switch oper {
			case "addx":
				wait += 2
				x += value
			case "noop":
				wait++
			default:
				log.Panic("You should not see this")
			}

			n++
		} else {
			wait--
		}

		cPos := ((cycle - 1) % 240) % 40
		if cPos+1 >= prevX && cPos-1 <= prevX {
			screen[(cycle-1)%240] = "#"
		} else {
			screen[(cycle-1)%240] = "."
		}

		cycle++
	}

	fmt.Println(screen)

	output := ""
	for i := 0; i < 6; i++ {
		output += fmt.Sprintln(strings.Join(screen[i*40:(i+1)*40], ""))
	}
	return output
}
