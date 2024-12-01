package day09

import "testing"
import _ "embed"

//go:embed example1
var example1Input string

//go:embed part1
var realInput string

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
			name:  "example",
			args:  args{example1Input},
			want:  114,
			want1: 0,
		},
		{
			name:  "real",
			args:  args{realInput},
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

func TestExtrapolate(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "",
			args:  args{[]int{0, 3, 6, 9, 12, 15}},
			want:  18,
			want1: 0,
		},
		{
			name:  "",
			args:  args{[]int{1, 3, 6, 10, 15, 21}},
			want:  28,
			want1: 0,
		},
		{
			name:  "example line 3",
			args:  args{[]int{10, 13, 16, 21, 30, 45}},
			want:  68,
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Extrapolate(tt.args.values)
			if got != tt.want {
				t.Errorf("Extrapolate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Extrapolate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
