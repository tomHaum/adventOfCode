package main

import "testing"
import _ "embed"

//go:embed part1
var part1Input string

//go:embed example
var exampleInput string

func Test_run(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name        string
		args        args
		wantAnswer1 int
		wantAnswer2 int
	}{
		{
			name: "example",
			args: args{
				input: exampleInput,
			},
			wantAnswer1: 4361,
			wantAnswer2: 467835,
		},
		{
			name:        "real",
			args:        args{part1Input},
			wantAnswer1: 537832, // not 536655
			wantAnswer2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAnswer1, gotAnswer2 := run(tt.args.input)
			if gotAnswer1 != tt.wantAnswer1 {
				t.Errorf("run() gotAnswer1 = %v, want %v", gotAnswer1, tt.wantAnswer1)
			}
			if gotAnswer2 != tt.wantAnswer2 {
				t.Errorf("run() gotAnswer2 = %v, want %v", gotAnswer2, tt.wantAnswer2)
			}
		})
	}
}
