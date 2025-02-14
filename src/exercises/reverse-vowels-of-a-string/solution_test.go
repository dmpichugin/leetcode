package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseVowels(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "abA",
			},
			want: "Aba",
		},
		{
			name: "case 2",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "case 3",
			args: args{
				s: "Abbo",
			},
			want: "obbA",
		},
		{
			name: "case 4",
			args: args{
				s: "AAbo",
			},
			want: "oAbA",
		},
		{
			name: "case ",
			args: args{
				s: "oaoa",
			},
			want: "aoao",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := reverseVowels(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
