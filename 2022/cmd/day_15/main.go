package main

import (
	"embed"
	"fmt"
	"log"
	"math"
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
	fmt.Printf("Task1: %q\n", Task1(string(input1), 2000000))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(string(input2)))
}

func Task1(input string, row int) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	sensors := make([]Point, len(lines))
	beacons := make([]Point, len(lines))

	maxWidth := 0
	minWidth := math.MaxInt
	maxDepth := 0
	minDepth := math.MaxInt

	for i, line := range lines {
		line = line[len("Sensor at "):]
		lineParts := strings.Split(line, ": closest beacon is at ")
		sensorInfo := strings.Split(lineParts[0], ", ")
		sensors[i] = NewPoint(sensorInfo[0][2:], sensorInfo[1][2:])
		if sensors[i].X > maxWidth {
			maxWidth = sensors[i].X
		}
		if sensors[i].X < minWidth {
			minWidth = sensors[i].X
		}
		if sensors[i].Y > maxDepth {
			maxDepth = sensors[i].Y
		}
		if sensors[i].Y < minDepth {
			minDepth = sensors[i].Y
		}
		beaconsInfo := strings.Split(lineParts[1], ", ")
		beacons[i] = NewPoint(beaconsInfo[0][2:], beaconsInfo[1][2:])
		if beacons[i].X > maxWidth {
			maxWidth = beacons[i].X
		}
		if beacons[i].X < minWidth {
			minWidth = beacons[i].X
		}
		if beacons[i].Y > maxDepth {
			maxDepth = beacons[i].Y
		}
		if beacons[i].Y < minDepth {
			minDepth = beacons[i].Y
		}
	}

	m := NewMap(maxWidth, minWidth, maxDepth, minDepth)
	m.SetSensors(sensors...)
	m.SetBeacons(beacons...)
	for _, sensor := range sensors {
		m.MarkSignal(sensor)
	}

	accum := 0
	fmt.Println(m.Tiles[row+m.yOffset])
	for _, elem := range m.Tiles[row+m.yOffset] {
		if elem == '#' {
			accum++
		}
	}

	return fmt.Sprint(accum)
}

func Task2(input string) string {
	return input
}

type Point struct {
	X, Y int
}

func NewPoint(x, y string) Point {
	p := Point{}
	var err error
	p.X, err = strconv.Atoi(x)
	if err != nil {
		log.Panic(err)
	}
	p.Y, err = strconv.Atoi(y)
	if err != nil {
		log.Panic(err)
	}
	return p
}

type Map struct {
	Tiles            [][]rune
	xOffset, yOffset int
}

func NewMap(maxWidth, minWidth, maxDepth, minDepth int) Map {
	m := Map{}
	if minWidth < 0 {
		m.xOffset = minWidth * -1
	}
	if minDepth < 0 {
		m.yOffset = minDepth * -1
	}

	m.Tiles = make([][]rune, maxDepth+m.yOffset+1)
	for i := range m.Tiles {
		m.Tiles[i] = make([]rune, maxWidth+m.xOffset+1)
	}

	return m
}

func (m Map) SetSensors(sensors ...Point) {
	for _, sensor := range sensors {
		m.Tiles[sensor.Y+m.yOffset][sensor.X+m.xOffset] = 'S'
	}
}

func (m Map) SetBeacons(beacons ...Point) {
	for _, beacon := range beacons {
		m.Tiles[beacon.Y+m.yOffset][beacon.X+m.xOffset] = 'B'
	}
}

func (m Map) MarkSignal(sensor Point) {
	next, stop := m.next(sensor)
	var future []Point
	for {
		for _, p := range next {
			if m.Tiles[p.Y+m.yOffset][p.X+m.xOffset] == 0 {
				m.Tiles[p.Y+m.yOffset][p.X+m.xOffset] = '#'
			}
			points, beaconFound := m.next(p)
			if beaconFound {
				stop = true
			}
			future = append(future, points...)
		}
		next = future
		if stop {
			break
		}
	}
	for _, p := range next {
		if m.Tiles[p.Y+m.yOffset][p.X+m.xOffset] == 0 {
			m.Tiles[p.Y+m.yOffset][p.X+m.xOffset] = '#'
		}
		m.next(p)
	}
}

func (m Map) Display() {
	for i := range m.Tiles {
		for j := range m.Tiles[i] {
			fmt.Print(string(m.Tiles[i][j]))
		}
		fmt.Println()
	}
}

func (m Map) next(p Point) ([]Point, bool) {
	var next []Point
	exit := false

	if len(m.Tiles) > p.Y+m.yOffset+1 {
		switch m.Tiles[p.Y+m.yOffset+1][p.X+m.xOffset] {
		case 0, 'S', '#':
			next = append(next, Point{Y: p.Y + 1, X: p.X})
		case 'B':
			exit = true
		}
	}

	if len(m.Tiles[p.Y+m.yOffset]) > p.X+m.xOffset+1 {
		switch m.Tiles[p.Y+m.yOffset][p.X+m.xOffset+1] {
		case 0, 'S', '#':
			next = append(next, Point{Y: p.Y, X: p.X + 1})
		case 'B':
			exit = true
		}
	}

	if p.Y+m.yOffset-1 >= 0 {
		switch m.Tiles[p.Y+m.yOffset-1][p.X+m.xOffset] {
		case 0, 'S', '#':
			next = append(next, Point{Y: p.Y - 1, X: p.X})
		case 'B':
			exit = true
		}
	}

	if p.X+m.xOffset-1 >= 0 {
		switch m.Tiles[p.Y+m.yOffset][p.X+m.xOffset-1] {
		case 0, 'S', '#':
			next = append(next, Point{Y: p.Y, X: p.X - 1})
		case 'B':
			exit = true
		}
	}

	return next, exit
}
