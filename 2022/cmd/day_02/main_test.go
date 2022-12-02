package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := `A Y
B X
C Z`
	expected := "15"

	output := Task1(input)
	require.Equal(t, expected, output)
}

func TestTask2(t *testing.T) {
	t.Parallel()

	input := `A Y
B X
C Z`
	expected := "12"

	output := Task2(input)
	require.Equal(t, expected, output)
}
