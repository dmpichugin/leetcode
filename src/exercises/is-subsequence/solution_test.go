package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSubsequence(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "abc",
				t: "a",
			},
			want: false,
		},
		{
			name: "case 1",
			args: args{
				s: "abc",
				t: "ahbgdcccc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isSubsequence(tt.args.s, tt.args.t)
			require.Equal(t, tt.want, got)
		})
	}
}
