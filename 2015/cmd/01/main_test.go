package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "(())",
			want:  "0",
		},
		{
			input: "()()",
			want:  "0",
		},
		{
			input: "(((",
			want:  "3",
		},
		{
			input: "(()(()(",
			want:  "3",
		},
		{
			input: "))(((((",
			want:  "3",
		},
		{
			input: "())",
			want:  "-1",
		},
		{
			input: "))(",
			want:  "-1",
		},
		{
			input: ")))",
			want:  "-3",
		},
		{
			input: ")())())",
			want:  "-3",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Task1(tt.input); got != tt.want {
				t.Errorf("Task1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: ")",
			want:  "1",
		},
		{
			input: "()())",
			want:  "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Task2(tt.input); got != tt.want {
				t.Errorf("Task2() = %v, want %v", got, tt.want)
			}
		})
	}
}
