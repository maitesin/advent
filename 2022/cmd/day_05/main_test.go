package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	expected := "CMZ"

	output := Task1(input)
	require.Equal(t, expected, output)
}

func TestTask2(t *testing.T) {
	t.Parallel()

	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	expected := "MCD"

	output := Task2(input)
	require.Equal(t, expected, output)
}
