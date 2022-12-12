package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	t.Parallel()

	monkeys := []Monkey{
		{
			Items: []int{79, 98},
			Operation: func(v int, mod int) int {
				return (v * 19) % mod
			},
			Test: func(v int) int {
				if v%23 == 0 {
					return 2
				}
				return 3
			},
			Eval: func() int {
				return 23
			},
		},
		{
			Items: []int{54, 65, 75, 74},
			Operation: func(v int, mod int) int {
				return (v + 6) % mod
			},
			Test: func(v int) int {
				if v%19 == 0 {
					return 2
				}
				return 0
			},
			Eval: func() int {
				return 19
			},
		},
		{
			Items: []int{79, 60, 97},
			Operation: func(v int, mod int) int {
				return (v * v) % mod
			},
			Test: func(v int) int {
				if v%13 == 0 {
					return 1
				}
				return 3
			},
			Eval: func() int {
				return 13
			},
		},
		{
			Items: []int{74},
			Operation: func(v int, mod int) int {
				return (v + 3) % mod
			},
			Test: func(v int) int {
				if v%17 == 0 {
					return 0
				}
				return 1
			},
			Eval: func() int {
				return 17
			},
		},
	}
	expected := "10605"

	output := Task1(monkeys)
	require.Equal(t, expected, output)
}

func TestTask2(t *testing.T) {
	t.Parallel()

	monkeys := []Monkey{
		{
			Items: []int{79, 98},
			Operation: func(v int, mod int) int {
				return (v * 19) % mod
			},
			Test: func(v int) int {
				if v%23 == 0 {
					return 2
				}
				return 3
			},
			Eval: func() int {
				return 23
			},
		},
		{
			Items: []int{54, 65, 75, 74},
			Operation: func(v int, mod int) int {
				return (v + 6) % mod
			},
			Test: func(v int) int {
				if v%19 == 0 {
					return 2
				}
				return 0
			},
			Eval: func() int {
				return 19
			},
		},
		{
			Items: []int{79, 60, 97},
			Operation: func(v int, mod int) int {
				return (v * v) % mod
			},
			Test: func(v int) int {
				if v%13 == 0 {
					return 1
				}
				return 3
			},
			Eval: func() int {
				return 13
			},
		},
		{
			Items: []int{74},
			Operation: func(v int, mod int) int {
				return (v + 3) % mod
			},
			Test: func(v int) int {
				if v%17 == 0 {
					return 0
				}
				return 1
			},
			Eval: func() int {
				return 17
			},
		},
	}
	expected := "2713310158"

	output := Task2(monkeys)
	require.Equal(t, expected, output)
}
