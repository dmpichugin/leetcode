package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isIsomorphic(t *testing.T) {
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
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isIsomorphic(tt.args.s, tt.args.t)
			require.Equal(t, tt.want, got)
		})
	}
}
