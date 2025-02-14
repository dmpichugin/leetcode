package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTheDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "case1 ",
			args: args{
				s: "",
				t: "y",
			},
			want: byte('y'),
		},
		{
			name: "case 2",
			args: args{
				s: "abcd",
				t: "abycd",
			},
			want: byte('y'),
		},
		{
			name: "case 2",
			args: args{
				s: "abcd",
				t: "yabcd",
			},
			want: byte('y'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := findTheDifference(tt.args.s, tt.args.t)
			require.Equal(t, tt.want, got)
		})
	}
}
