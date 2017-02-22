package template

import (
	"reflect"
	"testing"
)

func Test_validateFn2IntArgs(t *testing.T) {
	type args struct {
		ttag string
		tag  string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{
		{
			name:  "valid",
			args:  args{"word", "word(0, 2)"},
			want:  []int{0, 2},
			want1: true,
		},
		{
			name:  "valid no space",
			args:  args{"word", "word(0,2)"},
			want:  []int{0, 2},
			want1: true,
		},
		{
			name:  "valid custom",
			args:  args{"words", "words(0, 2)"},
			want:  []int{0, 2},
			want1: true,
		},
		{
			name:  "invalid non digit",
			args:  args{"word", "word('asd', 2)"},
			want:  nil,
			want1: false,
		},
		{
			name:  "invalid tag",
			args:  args{"words", "word(0, 2)"},
			want:  nil,
			want1: false,
		},
		{
			name:  "invalid tag fn",
			args:  args{"word", "word(0, 2}"},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validateFn2IntArgs(tt.args.ttag, tt.args.tag)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateWord() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("validateWord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
