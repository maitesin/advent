package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`,
			want: "114",
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
			input: `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`,
			want: "2",
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
