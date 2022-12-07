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
	fmt.Printf("Task1: %q\n", Task1(string(input1)))

	input2, err := inputsFS.ReadFile("inputs/2.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Task2: %q\n", Task2(string(input2)))
}

func Task1(input string) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	rootDir := NewDir(nil)
	dir := rootDir
	for i := 0; i < len(lines); i++ {
		switch {
		case lines[i] == "$ cd /":
			dir = rootDir
		case lines[i] == "$ cd ..":
			dir = dir.ParentDir
		case strings.HasPrefix(lines[i], "$ cd "):
			dir = dir.SubDirs[lines[i][5:]]
		case lines[i] == "$ ls":
			for i < len(lines)-1 {
				i++
				if strings.HasPrefix(lines[i], "$ ") {
					i--
					break
				}
				parts := strings.Split(lines[i], " ")
				switch parts[0] {
				case "dir":
					dir.SubDirs[parts[1]] = NewDir(dir)
				default:
					size, err := strconv.Atoi(parts[0])
					if err != nil {
						log.Panic(err)
					}
					dir.Files = append(dir.Files, File{Name: parts[1], Size: size})
				}
			}
		}
	}

	found := map[*Dir]struct{}{}
	rootDir.AtMost(100000, found)
	accum := 0
	for d := range found {
		accum += d.Size()
	}

	return fmt.Sprint(accum)
}

func Task2(input string) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	rootDir := NewDir(nil)
	dir := rootDir
	for i := 0; i < len(lines); i++ {
		switch {
		case lines[i] == "$ cd /":
			dir = rootDir
		case lines[i] == "$ cd ..":
			dir = dir.ParentDir
		case strings.HasPrefix(lines[i], "$ cd "):
			dir = dir.SubDirs[lines[i][5:]]
		case lines[i] == "$ ls":
			for i < len(lines)-1 {
				i++
				if strings.HasPrefix(lines[i], "$ ") {
					i--
					break
				}
				parts := strings.Split(lines[i], " ")
				switch parts[0] {
				case "dir":
					dir.SubDirs[parts[1]] = NewDir(dir)
				default:
					size, err := strconv.Atoi(parts[0])
					if err != nil {
						log.Panic(err)
					}
					dir.Files = append(dir.Files, File{Name: parts[1], Size: size})
				}
			}
		}
	}

	max := 70000000
	needed := 30000000
	used := rootDir.Size()

	threshold := -1 * (max - needed - used)
	return fmt.Sprint(rootDir.NextBiggest(threshold, max))
}

type File struct {
	Name string
	Size int
}

type Dir struct {
	SubDirs   map[string]*Dir
	Files     []File
	ParentDir *Dir
}

func NewDir(parent *Dir) *Dir {
	d := Dir{}
	d.SubDirs = map[string]*Dir{}
	d.ParentDir = parent
	return &d
}

func (d *Dir) Size() int {
	subDirSizes := 0

	for _, dir := range d.SubDirs {
		subDirSizes += dir.Size()
	}

	filesSize := 0

	for _, f := range d.Files {
		filesSize += f.Size
	}

	return subDirSizes + filesSize
}

func (d *Dir) AtMost(threshold int, found map[*Dir]struct{}) {
	for _, sDir := range d.SubDirs {
		if sDir.Size() <= threshold {
			if _, ok := found[sDir]; !ok {
				found[sDir] = struct{}{}
			}
		}
		sDir.AtMost(threshold, found)
	}
}

func (d *Dir) NextBiggest(threshold int, current int) int {
	for _, sDir := range d.SubDirs {
		size := sDir.Size()
		if size > threshold {
			if size < current {
				current = size
			}
		}
		sub := sDir.NextBiggest(threshold, current)
		if sub < current {
			return sub
		}
	}

	return current
}
