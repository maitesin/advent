package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	input := ``
	expected := ""

	output := Task1(io.NopCloser(strings.NewReader(input)))
	require.Equal(t, expected, output)
}

func TestTask2(t *testing.T) {
	t.Parallel()

	input := ``
	expected := ""

	output := Task2(io.NopCloser(strings.NewReader(input)))
	require.Equal(t, expected, output)
}
