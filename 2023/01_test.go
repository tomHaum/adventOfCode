package adventofcode2023

import (
	_ "embed"
	"strings"
	"testing"
)

func Test_calibrationValue(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example line 1",
			args: args{
				line: "1abc2",
			},
			want: 12,
		},
		{
			name: "example line 2",
			args: args{
				line: "pqr3stu8vwx",
			},
			want: 38,
		},
		{
			name: "example line 3",
			args: args{
				line: "a1b2c3d4e5f",
			},
			want: 15,
		},
		{
			name: "example line 4",
			args: args{
				line: "treb7uchet",
			},
			want: 77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calibrationValue(tt.args.line); got != tt.want {
				t.Errorf("calibrationValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:embed day1-part1
var day1Part1InputRaw string
var day1Part1Input []string

var day1Part1ExampleInput = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}
var day1Part2ExampleInput = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func init() {
	day1Part1Input = strings.Split(day1Part1InputRaw, "\n")
	for i, line := range day1Part1Input {
		day1Part1Input[i] = strings.TrimSpace(line)
	}
}

func Test_day1Part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: day1Part1ExampleInput,
			},
			want: 142,
		},
		{
			name: "real",
			args: args{
				input: day1Part1Input,
			},
			want: 55488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1Part1(tt.args.input); got != tt.want {
				t.Errorf("day1Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calibrateValue2(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example line 1",
			args: args{"two1nine"},
			want: 29,
		},
		{
			name: "example line 1",
			args: args{"eightwothree"},
			want: 83,
		},
		{
			name: "example line 1",
			args: args{"abcone2threexyz"},
			want: 13,
		},
		{
			name: "example line 1",
			args: args{"xtwone3four"},
			want: 24,
		},
		{
			name: "example line 1",
			args: args{"4nineeightseven2"},
			want: 42,
		},
		{
			name: "example line 1",
			args: args{"zoneight234"},
			want: 14,
		},
		{
			name: "example line 1",
			args: args{"7pqrstsixteen"},
			want: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calibrateValue2(tt.args.line); got != tt.want {
				t.Errorf("calibrateValue2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day1part2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: day1Part2ExampleInput,
			},
			want: 281,
		},
		{
			name: "real",
			args: args{
				input: day1Part1Input,
			},
			want: 55614,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1part2(tt.args.input); got != tt.want {
				t.Errorf("day1part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
