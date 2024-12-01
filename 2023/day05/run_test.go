package day05

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed example
var exampleInput string

//go:embed part1input
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
			want:  35,
			want1: 0,
		},
		{
			name: "real",
			args: args{
				input: part1Input,
			},
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

func Test_getSeeds(t *testing.T) {
	type args struct {
		split []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example",
			args: args{
				split: []string{
					"seeds: 79 14 55 13",
				},
			},
			want: []int{79, 14, 55, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeeds(tt.args.split); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTransformation(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Transformation
	}{
		{
			name: "example 1",
			args: args{"50 98 2"},
			want: Transformation{
				SourceStart: 98,
				SourceEnd:   99,
				Delta:       -48,
			},
		},
		{
			name: "example 2",
			args: args{line: "52 50 48"},
			want: Transformation{
				SourceStart: 50,
				SourceEnd:   97,
				Delta:       2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTransformation(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTransformation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTransformer(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want Transformer
	}{
		{
			name: "example 1",
			args: args{
				input: []string{
					"seed-to-soil map:",
					"50 98 2",
					"52 50 48",
				},
			},
			want: Transformer{
				Source:      "seed",
				Destination: "soil",
				Tranformations: []Transformation{
					{
						SourceStart: 98,
						SourceEnd:   99,
						Delta:       -48,
					},
					{
						SourceStart: 50,
						SourceEnd:   97,
						Delta:       2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTransformer(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTransformer() = %v, want %v", got, tt.want)
			}
		})
	}
}
