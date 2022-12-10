package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	expected := "13"

	output := Task1(input)
	require.Equal(t, expected, output)
}

func TestTask2_1(t *testing.T) {
	t.Parallel()

	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	expected := "1"

	output := Task2(input)
	require.Equal(t, expected, output)
}

func TestTask2_2(t *testing.T) {
	t.Parallel()

	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	expected := "36"

	output := Task2(input)
	require.Equal(t, expected, output)
}
