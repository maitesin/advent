package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := `30373
25512
65332
33549
35390`
	expected := "21"

	output := Task1(input)
	require.Equal(t, expected, output)
}

func TestTask2(t *testing.T) {
	t.Parallel()

	input := `30373
25512
65332
33549
35390`
	expected := "8"

	output := Task2(input)
	require.Equal(t, expected, output)
}
