package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	straight bool
	limits   [2]Point
	points   []Point
}

func NewLine(s, e Point) Line {
	straight := s.x == e.x || s.y == e.y

	xMax := 0
	xMin := 0
	yMax := 0
	yMin := 0

	for _, point := range []Point{s, e} {
		if point.x < xMin {
			xMin = point.x
		}
		if point.x > xMax {
			xMax = point.x
		}
		if point.y < yMin {
			yMin = point.y
		}
		if point.y > yMax {
			yMax = point.y
		}
	}

	limits := [2]Point{{xMin, yMin}, {xMax, yMax}}

	var points []Point
	if straight {
		if s.x == e.x {
			for i := s.y; i <= e.y; i++ {
				points = append(points, Point{s.x, i})
			}
		} else {
			for i := s.x; i <= e.x; i++ {
				points = append(points, Point{i, s.y})
			}
		}
	} else {
		if s.y-e.y > 0 {
			offset := 0
			for i := s.x; i <= e.x; i++ {
				points = append(points, Point{s.x + offset, s.y - offset})
				offset++
			}
		} else {
			offset := 0
			for i := s.x; i <= e.x; i++ {
				points = append(points, Point{s.x + offset, s.y + offset})
				offset++
			}
		}
	}

	return Line{
		straight: straight,
		limits:   limits,
		points:   points,
	}
}

func (l *Line) inLine(p Point) bool {
	for _, point := range l.points {
		if point == p {
			return true
		}
	}
	return false
}

func findLimits(lines []Line) [2]Point {
	xMax := 0
	xMin := 0
	yMax := 0
	yMin := 0

	for _, line := range lines {
		if line.limits[0].x < xMin {
			xMin = line.limits[0].x
		}
		if line.limits[0].y < yMin {
			yMin = line.limits[0].y
		}
		if line.limits[1].x > xMax {
			xMax = line.limits[1].x
		}
		if line.limits[1].y > yMax {
			yMax = line.limits[1].y
		}
	}

	return [2]Point{{xMin, yMin}, {xMax + 1, yMax + 1}}
}

func main() {
	f1 := readInput()
	lines := listOfLines(f1)

	fmt.Printf("Answer #1: %d\n", overlaps(keepStraightLines(lines)))
	fmt.Printf("Answer #2: %d\n", overlaps(lines))
}

func overlaps(lines []Line) int {
	points := findLimits(lines)
	min, max := points[0], points[1]

	accum := 0
	for i := min.y; i < max.y; i++ {
		for j := min.x; j < max.x; j++ {
			covering := 0
			p := Point{j, i}
			for _, line := range lines {
				if line.inLine(p) {
					covering++
				}
			}
			if covering > 1 {
				accum++
			}
		}
	}

	return accum
}

func listOfLines(closer io.ReadCloser) []Line {
	buffer := bufio.NewReader(closer)

	var lines []Line

	for {
		line, _, err := buffer.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("%+v\n", err))
		}
		points := strings.Split(string(line), "->")
		p0 := parsePoint(points[0])
		p1 := parsePoint(points[1])

		if p0.x == p1.x {
			if p0.y > p1.y {
				p0, p1 = p1, p0
			}
		} else {
			if p0.x > p1.x {
				p0, p1 = p1, p0
			}
		}
		lines = append(lines, NewLine(p0, p1))
	}

	return lines
}

func keepStraightLines(lines []Line) []Line {
	var straightLines []Line

	for _, line := range lines {
		if line.straight {
			straightLines = append(straightLines, line)
		}
	}

	return straightLines
}

func parsePoint(point string) Point {
	parts := strings.Split(strings.TrimSpace(point), ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return Point{
		x: x,
		y: y,
	}
}

func readInput() io.ReadCloser {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	return file
}
