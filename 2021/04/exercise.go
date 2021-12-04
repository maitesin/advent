package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value  int
	marked bool
}

type Board struct {
	numbers [5][5]Number
	done    bool
}

func readBoard(reader *bufio.Reader) (Board, error) {
	var board Board

	lineStr, _, err := reader.ReadLine()
	if err != nil {
		return board, err
	}

	if len(lineStr) != 0 {
		return board, fmt.Errorf("not empty line found")
	}

	for i := range board.numbers {
		line, err := readLine(reader)
		if err != nil {
			return board, err
		}

		board.numbers[i] = line
	}

	return board, nil
}

func readLine(reader *bufio.Reader) ([5]Number, error) {
	var line [5]Number

	lineStr, _, err := reader.ReadLine()
	if err != nil {
		return line, err
	}

	lineStrParts := strings.Split(strings.TrimSpace(string(lineStr)), " ")
	var cleanedParts []string
	for i := range lineStrParts {
		if lineStrParts[i] != "" {
			cleanedParts = append(cleanedParts, lineStrParts[i])
		}
	}

	for i := range line {
		part, err := strconv.Atoi(cleanedParts[i])
		if err != nil {
			return line, err
		}

		line[i] = Number{value: part}
	}

	return line, nil
}

func (b *Board) Mark(value int) {
	for i := range b.numbers {
		for j := range b.numbers[i] {
			if b.numbers[i][j].value == value {
				b.numbers[i][j].marked = true
			}
		}
	}
}

func (b *Board) Check() bool {
	for _, line := range b.numbers {
		marked := 0
		for _, number := range line {
			if number.marked {
				marked++
			}
		}
		if marked == 5 {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		marked := 0
		for j := 0; j < 5; j++ {
			if b.numbers[j][i].marked {
				marked++
			}
		}
		if marked == 5 {
			return true
		}
	}

	return false
}

func (b *Board) AddUpUnmarked() int {
	accum := 0
	for _, line := range b.numbers {
		for _, number := range line {
			if !number.marked {
				accum += number.value
			}
		}
	}
	return accum
}

func removeDoneBoards(boards []Board) []Board {
	var undoneBoards []Board
	for _, board := range boards {
		if !board.done {
			undoneBoards = append(undoneBoards, board)
		}
	}

	return undoneBoards
}

func main() {
	f1 := readInput()
	draws, boards := listOfDrawsAndBoards(f1)

	first := true
	for _, draw := range draws {
		for i := range boards {
			boards[i].Mark(draw)
			if boards[i].Check() {
				if first {
					fmt.Printf("Answer #1: %d\n", draw*boards[i].AddUpUnmarked())
					first = false
				} else {
					if len(boards) == 1 {
						fmt.Printf("Answer #2: %d\n", draw*boards[0].AddUpUnmarked())
						return
					}
				}
				boards[i].done = true
			}
		}
		boards = removeDoneBoards(boards)
	}
}

func listOfDrawsAndBoards(closer io.ReadCloser) ([]int, []Board) {
	buffer := bufio.NewReader(closer)

	drawsStr, _, err := buffer.ReadLine()
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	drawsStrParts := strings.Split(strings.TrimSpace(string(drawsStr)), ",")
	draws := make([]int, len(drawsStrParts))
	for i, part := range drawsStrParts {
		draw, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("%+v\n", err))
		}
		draws[i] = draw
	}

	var boards []Board
	for {
		board, err := readBoard(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("%+v\n", err))
		}

		boards = append(boards, board)
	}

	return draws, boards
}

func readInput() io.ReadCloser {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	return file
}
