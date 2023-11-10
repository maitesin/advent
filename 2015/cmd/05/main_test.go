package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "ugknbfddgicrmopn",
			want:  "1",
		},
		{
			input: "aaa",
			want:  "1",
		},
		{
			input: "jchzalrnumimnmhp",
			want:  "0",
		},
		{
			input: "haegwjzuvuyypxyu",
			want:  "0",
		},
		{
			input: "dvszwmarrgswjxmb",
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
			input: "qjhvhtzxzqqjkmpb",
			want:  "1",
		},
		{
			input: "xxyxx",
			want:  "1",
		},
		{
			input: "uurcxstgmygtbstg",
			want:  "0",
		},
		{
			input: "ieodomkazucvgmuy",
			want:  "0",
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
