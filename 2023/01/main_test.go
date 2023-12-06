package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "abcdef",
			want:  "609043",
		},
		{
			input: "pqrstuv",
			want:  "1048970",
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
			input: "^v",
			want:  "3",
		},
		{
			input: "^>v<",
			want:  "3",
		},
		{
			input: "^v^v^v^v^v",
			want:  "11",
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
