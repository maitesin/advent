package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1_1(t *testing.T) {
	t.Parallel()

	input := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	expected := "7"

	output := Task1(input)
	require.Equal(t, expected, output)
}
func TestTask1_2(t *testing.T) {
	t.Parallel()

	input := `bvwbjplbgvbhsrlpgdmjqwftvncz`
	expected := "5"

	output := Task1(input)
	require.Equal(t, expected, output)
}
func TestTask1_3(t *testing.T) {
	t.Parallel()

	input := `nppdvjthqldpwncqszvftbrmjlhg`
	expected := "6"

	output := Task1(input)
	require.Equal(t, expected, output)
}
func TestTask1_4(t *testing.T) {
	t.Parallel()

	input := `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`
	expected := "10"

	output := Task1(input)
	require.Equal(t, expected, output)
}
func TestTask1_5(t *testing.T) {
	t.Parallel()

	input := `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`
	expected := "11"

	output := Task1(input)
	require.Equal(t, expected, output)
}

func TestTask2_1(t *testing.T) {
	t.Parallel()

	input := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	expected := "19"

	output := Task2(input)
	require.Equal(t, expected, output)
}

func TestTask2_2(t *testing.T) {
	t.Parallel()

	input := `bvwbjplbgvbhsrlpgdmjqwftvncz`
	expected := "23"

	output := Task2(input)
	require.Equal(t, expected, output)
}

func TestTask2_3(t *testing.T) {
	t.Parallel()

	input := `nppdvjthqldpwncqszvftbrmjlhg`
	expected := "23"

	output := Task2(input)
	require.Equal(t, expected, output)
}

func TestTask2_4(t *testing.T) {
	t.Parallel()

	input := `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`
	expected := "29"

	output := Task2(input)
	require.Equal(t, expected, output)
}

func TestTask2_5(t *testing.T) {
	t.Parallel()

	input := `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`
	expected := "26"

	output := Task2(input)
	require.Equal(t, expected, output)
}
