package day06

import "testing"

func TestRace_Solutions(t *testing.T) {
	type fields struct {
		Time     int
		Distance int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "example 1",
			fields: fields{
				Time:     7,
				Distance: 9,
			},
			want: 4,
		},
		{
			name: "example 2",
			fields: fields{
				Time:     15,
				Distance: 40,
			},
			want: 8,
		},
		{
			name: "example 3",
			fields: fields{
				Time:     30,
				Distance: 200,
			},
			want: 9,
		},
		{
			name: "part 2",
			fields: fields{
				Time:     42899189,
				Distance: 308117012911467,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Race{
				Time:     tt.fields.Time,
				Distance: tt.fields.Distance,
			}
			if got := r.Solutions(); got != tt.want {
				t.Errorf("Solutions() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			name: "example",
			args: args{
				input: `Time:      7  15   30
Distance:  9  40  200`,
			},
			want:  288,
			want1: 0,
		},
		{
			name: "real",
			args: args{
				input: `Time:        42     89     91     89
Distance:   308   1170   1291   1467`,
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
