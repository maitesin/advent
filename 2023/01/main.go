package main

import (
	"crypto/md5"
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
	count := 1
	for ; ; count++ {
		hash := generateMD5(fmt.Sprintf("%s%d", input, count))
		if strings.HasPrefix(hash, "00000") {
			return fmt.Sprint(count)
		}
	}
}

func generateMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func Task2(input string) string {
	count := 1
	for ; ; count++ {
		hash := generateMD5(fmt.Sprintf("%s%d", input, count))
		if strings.HasPrefix(hash, "000000") {
			return fmt.Sprint(count)
		}
	}
}
