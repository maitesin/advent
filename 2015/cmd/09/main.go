package main

import (
	"embed"
	"fmt"
	"log"
	"math"
	"slices"
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

	//input2, err := inputsFS.ReadFile("inputs/2.txt")
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Printf("Task2: %q\n", Task2(strings.TrimSpace(string(input2))))
}

type Map struct {
	Cities map[string]map[string]int
}

func NewMap() Map {
	return Map{
		Cities: map[string]map[string]int{},
	}
}

func (m *Map) Display() {
	for k, v := range m.Cities {
		fmt.Printf("%s:\n", k)
		for kk, vv := range v {
			fmt.Printf("%s: %d\n", kk, vv)
		}
		fmt.Println()
	}
}

func (m *Map) FindShortedDistanceToNotAlreadyUsed(city string, visited map[string]struct{}) (string, int) {
	distance := math.MaxInt
	nextCity := ""
	for k, v := range m.Cities[city] {
		if _, ok := visited[k]; !ok {
			if v < distance {
				distance = v
				nextCity = k
			}
		}
	}
	return nextCity, distance
}

func Task1(input string) string {
	m := parseMap(input)
	cities := make([]string, 0, len(m.Cities))
	for k, _ := range m.Cities {
		cities = append(cities, k)
	}
	visited := make([]string, 0, len(m.Cities))
	visitedFast := make(map[string]struct{}, len(m.Cities))
	visited = append(visited, cities[0])
	visitedFast[cities[0]] = struct{}{}
	cities = cities[1:]
	distance := 0

	m.Display()

	for {
		if len(cities) == 0 {
			break
		}
		next1, distance1 := m.FindShortedDistanceToNotAlreadyUsed(visited[0], visitedFast)
		next2, distance2 := m.FindShortedDistanceToNotAlreadyUsed(visited[len(visited)-1], visitedFast)
		if distance1 < distance2 {
			visited = append([]string{next1}, visited...)
			visitedFast[next1] = struct{}{}
			distance += distance1
			cities = slices.DeleteFunc(cities, func(s string) bool {
				return s == next1
			})
		} else {
			visited = append(visited, next2)
			visitedFast[next2] = struct{}{}
			distance += distance2
			cities = slices.DeleteFunc(cities, func(s string) bool {
				return s == next2
			})
		}
	}
	fmt.Printf("%#v\n", visited)
	return fmt.Sprint(distance)
}

func parseMap(input string) Map {
	m := NewMap()
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		city1, city2, distance := parseLine(line)
		if _, ok := m.Cities[city1]; !ok {
			m.Cities[city1] = map[string]int{}
		}
		m.Cities[city1][city2] = distance
		if _, ok := m.Cities[city2]; !ok {
			m.Cities[city2] = map[string]int{}
		}
		m.Cities[city2][city1] = distance
	}

	return m
}

func parseLine(line string) (string, string, int) {
	sections := strings.Split(line, " = ")
	distance, err := strconv.Atoi(sections[1])
	if err != nil {
		panic(err)
	}
	cities := strings.Split(sections[0], " to ")
	return cities[0], cities[1], distance
}

func Task2(input string) string {
	return input
}
