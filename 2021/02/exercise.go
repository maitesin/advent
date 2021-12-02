package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type action struct {
	direction string
	length int
}

func main() {
	f1 := readInput()
	actions := listOfActions(f1)

	x := 0
	y := 0

	for _, action := range actions {
		switch action.direction {
		case "forward":
			x += action.length
		case "down":
			y += action.length
		case "up":
			y -= action.length
		default:
			panic(fmt.Sprintf("Invalid direction found in action %q", action.direction))
		}
	}

	fmt.Printf("Answer #1: %d\n", x * y)

	x = 0
	y = 0
	aim := 0

	for _, action := range actions {
		switch action.direction {
		case "forward":
			x += action.length
			y += action.length * aim
		case "down":
			aim += action.length
		case "up":
			aim -= action.length
		default:
			panic(fmt.Sprintf("Invalid direction found in action %q", action.direction))
		}
	}

	fmt.Printf("Answer #2: %d\n", x * y)
}

func listOfActions(closer io.ReadCloser) []action {
	buffer := bufio.NewReader(closer)

	var values []action
	for {
		actionStr, _, err := buffer.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("%+v\n", err))
		}
		actionParts := strings.Split(string(actionStr), " ")
		if len(actionParts) != 2 {
			panic(fmt.Sprintf("Invalid number of parts in the action %q", actionStr))
		}
		direction := actionParts[0]
		valueInt, err := strconv.Atoi(strings.TrimSpace(string(actionParts[1])))
		if err != nil {
			panic(fmt.Sprintf("%+v\n", err))
		}
		values = append(values, action{direction: direction, length: valueInt		})
	}
	return values
}


func readInput() io.ReadCloser {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	return file
}