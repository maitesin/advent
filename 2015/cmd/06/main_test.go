package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "turn on 0,0 through 999,999",
			want:  "1000000",
		},
		{
			input: "toggle 0,0 through 999,0",
			want:  "1000",
		},
		{
			input: "turn off 499,499 through 500,500",
			want:  "0",
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
			input: "turn on 0,0 through 0,0",
			want:  "1",
		},
		{
			input: "toggle 0,0 through 999,999",
			want:  "2000000",
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
