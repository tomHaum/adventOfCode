package day1

import "testing"
import _ "embed"

//go:embed example-part1
var example1Input string

//go:embed part1
var part1Input string

func Test_run(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want1 int
		want2 int
	}{
		{
			name:  "example",
			args:  args{input: example1Input},
			want1: 24000,
			want2: 45000,
		},
		{
			name:  "real",
			args:  args{input: part1Input},
			want1: 67622,
			want2: 201491,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPart1, gotPart2 := run(tt.args.input)
			if gotPart1 != tt.want1 {
				t.Errorf("run() = %v, want %v", gotPart1, tt.want1)
			}
			if gotPart2 != tt.want2 {
				t.Errorf("run() = %v, want %v", gotPart2, tt.want2)
			}
		})
	}
}
