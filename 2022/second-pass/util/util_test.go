package util

import (
	"reflect"
	"testing"
)

func Test_splitLines(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "windows",
			args: args{str: "line 1\r\nline2\r\n\r\nline4"},
			want: []string{"line 1", "line2", "", "line4"},
		},
		{
			name: "linux",
			args: args{str: "line 1\nline2\n\nline4"},
			want: []string{"line 1", "line2", "", "line4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLines(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
