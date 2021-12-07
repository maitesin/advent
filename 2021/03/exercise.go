package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f1 := readInput()
	values := listOfNumbers(f1)

	gamma := ""
	epsilon := ""
	for i := range values[0] {
		zeros, ones := countZerosAndOnesInPosition(values, i)
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer #1: %d\n", gammaDec * epsilonDec)

	oxygen := filterValuesByPattern(values[:], '1', '0')
	oxygenDec, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		panic(err)
	}

	co2 := filterValuesByPattern(values[:], '0', '1')
	co2Dec, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer #2: %d\n", oxygenDec * co2Dec)

}

func filterValuesByPattern(values []string, def, alt rune) string {
	l := len(values[0])
	for i := 0; i < l; i++ {
		zeros, ones := countZerosAndOnesInPosition(values, i)
		char := def
		if zeros > ones {
			char = alt
		}
		values = filterValuesByCharacterAndPosition(values, char, i)
		if len(values) == 1 {
			return values[0]
		}
	}
	return values[0]
}


func filterValuesByCharacterAndPosition(values []string, c rune, pos int) []string {
	var remaining []string
	for _, value := range values {
		if []rune(value)[pos] == c {
			remaining = append(remaining, value)
		}
	}

	return remaining
}

func countZerosAndOnesInPosition(values[]string, pos int) (int, int) {
	zeros := 0
	ones := 0
	for _, value := range values {
		switch value[pos] {
		case '0':
			zeros++
		case '1':
			ones++
		default:
			panic(fmt.Sprintf("invalid digit in position %d in value %q", pos, value))
		}
	}
	return zeros, ones
}

func listOfNumbers(closer io.ReadCloser) []string {
	buffer := bufio.NewReader(closer)

	var values []string
	for {
		value, _, err := buffer.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("%+v\n", err))
		}
		values = append(values, strings.TrimSpace(string(value)))
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