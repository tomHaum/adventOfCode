package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed day1.real.txt
var day1 string

//go:embed day1.sample.txt
var day1sample string

func Test_parse(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "sample ",
			args: args{
				data: day1sample,
			},
			want:  []int{3, 4, 2, 1, 3, 3},
			want1: []int{4, 3, 5, 3, 9, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDay1(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "sample",
			args: args{
				data: day1sample,
			},
			want:  11,
			want1: 31,
		},
		{
			name: "real",
			args: args{
				data: day1,
			},
			want:  3569916,
			want1: 26407426,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, right := parse(tt.args.data)
			got, got1 := Day1(left, right)
			if got != tt.want {
				t.Errorf("Day1() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Day1() got1 = %v, want %v", got1, tt.want1)
			}

			t.Logf("%v. Part1: %v. Part2: %v\n", tt.name, got, got1)
		})
	}
}
