package day04

import "testing"
import _ "embed"

//go:embed example
var exampleInput string

//go:embed part1
var part1Input string

func Test_run(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "example",
			args: args{
				input: exampleInput,
			},
			want:  13,
			want1: 30,
		},
		{
			name: "real",
			args: args{
				input: part1Input,
			},
			want:  0,
			want1: 0, // not 3390317
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := run(tt.args.input)
			if got != tt.want {
				t.Errorf("run() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("run() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
