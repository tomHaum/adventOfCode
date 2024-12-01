package day07

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestParseHand(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Hand
	}{
		{
			name: "exmaple 1",
			args: args{"32T3K 765"},
			want: Hand{
				Cards: []Card("32T3K"),
				Bet:   765,
				Type:  Type1Pair,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseHand(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:embed example1
var example1Input string

//go:embed part1
var part1Input string

func TestRun(t *testing.T) {
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
			want:  6440,
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
			got, got1 := Run(tt.args.input)
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
