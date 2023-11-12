package main

import (
	"embed"
	"fmt"
	"log"
	"reflect"
	"sort"
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

type status struct {
	Value uint16
	Set   bool
}

type board struct {
	Values       map[string]status
	Instructions []instruction
}

func newBoard() board {
	return board{
		Values: map[string]status{},
	}
}

func (b *board) load(inst []instruction) {
	b.Instructions = inst
}

func (b *board) run() {
	for _, inst := range b.Instructions {
		switch inst.Operation {
		case "and":
			s1 := b.Values[inst.Sources[0]]
			s2 := status{
				Value: inst.Value,
				Set:   true,
			}
			if len(inst.Sources) > 1 {
				s2 = b.Values[inst.Sources[1]]
			}
			if s1.Set && s2.Set {
				b.Values[inst.Destination] = status{Value: s1.Value & s2.Value, Set: true}
			}
		case "or":
			s1 := b.Values[inst.Sources[0]]
			s2 := status{
				Value: inst.Value,
				Set:   true,
			}
			if len(inst.Sources) > 1 {
				s2 = b.Values[inst.Sources[1]]
			}
			if s1.Set && s2.Set {
				b.Values[inst.Destination] = status{Value: s1.Value | s2.Value, Set: true}
			}
		case "ls":
			s1 := b.Values[inst.Sources[0]]
			s2 := inst.Value
			if s1.Set {
				b.Values[inst.Destination] = status{Value: s1.Value << s2, Set: true}
			}
		case "rs":
			s1 := b.Values[inst.Sources[0]]
			s2 := inst.Value
			if s1.Set {
				b.Values[inst.Destination] = status{Value: s1.Value >> s2, Set: true}
			}
		case "not":
			s1 := b.Values[inst.Sources[0]]
			if s1.Set {
				b.Values[inst.Destination] = status{Value: ^s1.Value, Set: true}
			}
		case "value":
			if len(inst.Sources) > 0 {
				if b.Values[inst.Sources[0]].Set {
					b.Values[inst.Destination] = b.Values[inst.Sources[0]]
				}
			} else {
				b.Values[inst.Destination] = status{Value: inst.Value, Set: true}
			}
		default:
			panic("You should not see this")
		}
	}
}

func (b *board) display() string {
	keys := make([]string, 0, len(b.Values))
	for k := range b.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	output := strings.Builder{}
	for _, k := range keys {
		output.WriteString(fmt.Sprintf("%s: %d\n", k, b.Values[k].Value))
	}
	return output.String()
}

type instruction struct {
	Sources     []string
	Value       uint16
	Operation   string
	Destination string
}

func parseInstructions(input string) []instruction {
	parts := strings.Split(input, "\n")
	var instructions []instruction
	for _, part := range parts {
		sections := strings.Split(part, " -> ")
		inst := instruction{Destination: sections[1]}
		switch {
		case strings.Contains(sections[0], "AND"):
			opers := strings.Split(sections[0], " ")
			inst.Sources = []string{opers[0], opers[2]}
			v, err := strconv.Atoi(opers[0])
			if err != nil {
				inst.Sources = []string{opers[0], opers[2]}
			} else {
				inst.Sources = []string{opers[2]}
				inst.Value = uint16(v)
			}
			inst.Operation = "and"
		case strings.Contains(sections[0], "OR"):
			opers := strings.Split(sections[0], " ")
			v, err := strconv.Atoi(opers[0])
			if err != nil {
				inst.Sources = []string{opers[0], opers[2]}
			} else {
				inst.Sources = []string{opers[2]}
				inst.Value = uint16(v)
			}
			inst.Operation = "or"
		case strings.Contains(sections[0], "LSHIFT"):
			opers := strings.Split(sections[0], " ")
			inst.Sources = []string{opers[0]}
			inst.Operation = "ls"
			v, err := strconv.Atoi(opers[2])
			if err != nil {
				panic(err)
			}
			inst.Value = uint16(v)
		case strings.Contains(sections[0], "RSHIFT"):
			opers := strings.Split(sections[0], " ")
			inst.Sources = []string{opers[0]}
			inst.Operation = "rs"
			v, err := strconv.Atoi(opers[2])
			if err != nil {
				panic(err)
			}
			inst.Value = uint16(v)
		case strings.HasPrefix(sections[0], "NOT"):
			opers := strings.Split(sections[0], " ")
			v, err := strconv.Atoi(opers[1])
			if err != nil {
				inst.Sources = []string{opers[1]}
			} else {
				inst.Value = uint16(v)
			}
			inst.Operation = "not"
		default:
			v, err := strconv.Atoi(sections[0])
			if err != nil {
				inst.Sources = []string{sections[0]}
			} else {
				inst.Value = uint16(v)
			}
			inst.Operation = "value"
		}
		instructions = append(instructions, inst)
	}

	return instructions
}

func copyValues(original map[string]status) map[string]status {
	m := map[string]status{}
	for k, v := range original {
		m[k] = v
	}
	return m
}

func Task1(input string) string {
	instructions := parseInstructions(input)
	b := newBoard()
	b.load(instructions)
	for {
		previous := copyValues(b.Values)
		b.run()
		if reflect.DeepEqual(previous, b.Values) {
			break
		}
	}
	return b.display()
}

func Task2(input string) string {
	instructions := parseInstructions(input)
	b := newBoard()
	b.load(instructions)
	for {
		previous := copyValues(b.Values)
		b.run()
		if reflect.DeepEqual(previous, b.Values) {
			break
		}
	}
	return b.display()
}
