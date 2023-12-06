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
	lines := strings.Split(input, "\n")
	accum := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		b, e := findFirstAndLastDigits(line)
		v, err := strconv.Atoi(fmt.Sprintf("%d%d", b, e))
		if err != nil {
			panic(err)
		}
		accum += v
	}

	return fmt.Sprint(accum)
}

func findFirstAndLastDigits(line string) (int, int) {
	return findFirstDigit(line), findLastDigit(line)
}

func findFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			return calculateNumber(rune(line[i]))
		}
	}
	panic(fmt.Sprintf("Find first: %q", line))
}

func findLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return calculateNumber(rune(line[i]))
		}
	}
	panic(fmt.Sprintf("Find last: %q", line))
}

func calculateNumber(c rune) int {
	switch c {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	panic("Never! Switch")
}

func Task2(input string) string {
	lines := strings.Split(input, "\n")
	accum := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		b, e := findFirstAndLastNumbers(line)
		v, err := strconv.Atoi(fmt.Sprintf("%d%d", b, e))
		if err != nil {
			panic(err)
		}
		accum += v
	}

	return fmt.Sprint(accum)
}

func findFirstAndLastNumbers(line string) (int, int) {
	return findFirstNumber(line), findLastNumber(line)
}

func findFirstNumber(line string) int {
	for i := 0; i < len(line); i++ {
		start := startsWithNumber(line[i:])
		if line[i] >= '0' && line[i] <= '9' {
			return calculateNumber(rune(line[i]))
		} else if start != -1 {
			return start
		}
	}
	panic(fmt.Sprintf("Find first: %q", line))
}

func startsWithNumber(line string) int {
	switch {
	case strings.HasPrefix(line, "zero"):
		return 0
	case strings.HasPrefix(line, "one"):
		return 1
	case strings.HasPrefix(line, "two"):
		return 2
	case strings.HasPrefix(line, "three"):
		return 3
	case strings.HasPrefix(line, "four"):
		return 4
	case strings.HasPrefix(line, "five"):
		return 5
	case strings.HasPrefix(line, "six"):
		return 6
	case strings.HasPrefix(line, "seven"):
		return 7
	case strings.HasPrefix(line, "eight"):
		return 8
	case strings.HasPrefix(line, "nine"):
		return 9
	default:
		return -1
	}
}

func findLastNumber(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		ends := endsWithNumber(line[:i+1])
		if line[i] >= '0' && line[i] <= '9' {
			return calculateNumber(rune(line[i]))
		} else if ends != -1 {
			return ends
		}
	}
	panic(fmt.Sprintf("Find last: %q", line))
}

func endsWithNumber(line string) int {
	line = reverse(line)
	switch {
	case strings.HasPrefix(line, "orez"):
		return 0
	case strings.HasPrefix(line, "eno"):
		return 1
	case strings.HasPrefix(line, "owt"):
		return 2
	case strings.HasPrefix(line, "eerht"):
		return 3
	case strings.HasPrefix(line, "ruof"):
		return 4
	case strings.HasPrefix(line, "evif"):
		return 5
	case strings.HasPrefix(line, "xis"):
		return 6
	case strings.HasPrefix(line, "neves"):
		return 7
	case strings.HasPrefix(line, "thgie"):
		return 8
	case strings.HasPrefix(line, "enin"):
		return 9
	default:
		return -1
	}
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
