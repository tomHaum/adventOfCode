package day02

import (
	_ "embed"
	"reflect"
	"testing"
)

func Test_getId(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameId(tt.args.line); got != tt.want {
				t.Errorf("getGameId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGameSet(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want gameSet
	}{
		{
			name: "",
			args: args{"3 blue, 4 red"},
			want: gameSet{
				red:   4,
				blue:  3,
				green: 0,
			},
		},
		{
			name: "",
			args: args{"2 blue, 1 red, 2 green"},
			want: gameSet{
				red:   1,
				blue:  2,
				green: 2,
			},
		},
		{
			name: "",
			args: args{"2 green"},
			want: gameSet{green: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameSet(tt.args.txt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGameSets(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []gameSet
	}{
		{
			name: "example: game 1",
			args: args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			want: []gameSet{
				{red: 4, blue: 3, green: 0},
				{red: 1, blue: 6, green: 2},
				{red: 0, blue: 0, green: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameSets(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want game
	}{
		{
			name: "example: game 1",
			args: args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			want: game{
				id: 1,
				sets: []gameSet{
					{red: 4, blue: 3, green: 0},
					{red: 1, blue: 6, green: 2},
					{red: 0, blue: 0, green: 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGame(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:embed example-part1
var examplePart1Input string

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
			name: "example",
			args: args{
				input: examplePart1Input,
			},
			want1: 8,
			want2: 2286,
		},
		{
			name: "real",
			args: args{
				input: part1Input,
			},
			want1: 2239,
			want2: 83435,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := run(tt.args.input)
			if got1 != tt.want1 {
				t.Errorf("run() = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("run() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
