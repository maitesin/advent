package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	expected := "31"

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
