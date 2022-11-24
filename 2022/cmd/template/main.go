package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	input1 := ""
	fmt.Printf("Task1: %q\n", Task1(io.NopCloser(strings.NewReader(input1))))

	input2 := ""
	fmt.Printf("Task2: %q\n", Task2(io.NopCloser(strings.NewReader(input2))))
}

func Task1(rc io.ReadCloser) string {
	defer rc.Close()

	return ""
}

func Task2(rc io.ReadCloser) string {
	defer rc.Close()

	return ""
}
