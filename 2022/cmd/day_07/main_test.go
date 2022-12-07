package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
	expected := "95437"

	output := Task1(input)
	require.Equal(t, expected, output)
}

func TestTask2(t *testing.T) {
	t.Parallel()

	input := ``
	expected := ""

	output := Task2(input)
	require.Equal(t, expected, output)
}
