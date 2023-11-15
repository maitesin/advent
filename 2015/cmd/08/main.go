package main

import (
	"embed"
	"encoding/hex"
	"fmt"
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
	fmt.Printf("Task1: %q\n", Task1(strings.TrimSpace(string(input1))))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(strings.TrimSpace(string(input2))))
}

func Task1(input string) string {
	lines := strings.Split(input, "\n")
	charCount := 0
	decodedCount := 0
	for _, line := range lines {
		charCount += len(line)
		decoded := decodeLine(line)
		decodedCount += len(decoded)
	}
	return fmt.Sprint(charCount - decodedCount)
}
func decodeLine(line string) string {
	n := strings.Builder{}
	if line[0] == line[len(line)-1] && line[0] == '"' {
		line = line[1 : len(line)-1]
	}
	for i := 0; i < len(line); i++ {
		if line[i] == '\\' {
			i++
			if i >= len(line) {
				break
			}
			switch line[i] {
			case '\\', '"':
				n.WriteByte(line[i])
			case 'x':
				i++
				c, err := hex.DecodeString(line[i : i+2])
				if err != nil {
					panic(err)
				}
				n.Write(c)
				i++
			}
		} else {
			n.WriteByte(line[i])
		}
	}
	return n.String()
}

func Task2(input string) string {
	lines := strings.Split(input, "\n")
	charCount := 0
	encodedCount := 0
	for _, line := range lines {
		charCount += len(line)
		encoded := encodeLine(line)
		encodedCount += len(encoded)
		fmt.Printf("%s - %s; %d - %d\n", line, encoded, len(line), len(encoded))
	}
	return fmt.Sprint(encodedCount - charCount)
}

func encodeLine(input string) string {
	n := strings.Builder{}
	n.WriteRune('"')

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '"':
			n.WriteString("\\\"")
		case '\\':
			n.WriteString("\\\\")
		default:
			n.WriteByte(input[i])
		}
	}

	n.WriteRune('"')
	return n.String()
}
