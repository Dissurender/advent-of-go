package utils

import (
	"reflect"
	"testing"
)

func Test_sliceMaker(t *testing.T) {
	type args struct {
		str       string
		seperator string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should return 'hello', 'world'",
			args: args{str: "hello world", seperator: " "},
			want: []string{"hello", "world"},
		},
		{
			name: "should return '1', 'red'",
			args: args{str: "12 blue, 7 red, 12 green", seperator: ","},
			want: []string{"12 blue", "7 red", "12 green"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceMaker(tt.args.str, tt.args.seperator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceMaker() = %v, want %v", got, tt.want)
			}
		})
	}
}
