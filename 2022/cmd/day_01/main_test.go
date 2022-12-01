package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	expected := "24000"

	output := Task1(input)
	require.Equal(t, expected, output)
}

func TestTask2(t *testing.T) {
	t.Parallel()

	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	expected := "45000"

	output := Task2(input)
	require.Equal(t, expected, output)
}
