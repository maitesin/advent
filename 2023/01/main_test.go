package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: "142",
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
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: "281",
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
