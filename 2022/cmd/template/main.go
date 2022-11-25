package main

import (
	"embed"
	"fmt"
	"io"
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
	fmt.Printf("Task1: %q\n", Task1(io.NopCloser(strings.NewReader(string(input1)))))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(io.NopCloser(strings.NewReader(string(input2)))))
}

func Task1(rc io.ReadCloser) string {
	defer rc.Close()

	return ""
}

func Task2(rc io.ReadCloser) string {
	defer rc.Close()

	return ""
}
