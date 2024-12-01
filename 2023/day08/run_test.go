package day08

import (
	_ "embed"
	"testing"
)

//go:embed example1
var example1Input string

//go:embed example2
var example2Input string

//go:embed example3
var example3Input string

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
			name:  "example 1",
			args:  args{example1Input},
			want:  2,
			want1: 0,
		},
		{
			name:  "example 2",
			args:  args{example2Input},
			want:  6,
			want1: 0,
		},
		{
			name:  "example 3",
			args:  args{example3Input},
			want:  6,
			want1: 0,
		},
		{
			name:  "real",
			args:  args{part1Input},
			want:  0,
			want1: 0,
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

func TestLCM(t *testing.T) {
	type args struct {
		a        int
		b        int
		integers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(13201,
				14429,
				16271,
				20569,
				24253,
				21797); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}
