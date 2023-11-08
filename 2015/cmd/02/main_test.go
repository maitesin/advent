package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2x3x4",
			want:  "58",
		},
		{
			input: "1x1x10",
			want:  "43",
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
			input: "2x3x4",
			want:  "34",
		},
		{
			input: "1x1x10",
			want:  "14",
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
