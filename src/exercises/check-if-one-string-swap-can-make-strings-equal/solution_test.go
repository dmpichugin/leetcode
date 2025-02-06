package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAreAlmostEqual(t *testing.T) {
	t.Parallel()

	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				s1: "bank",
				s2: "kanb",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s1: "attack",
				s2: "defend",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				s1: "kelb",
				s2: "kelb",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := areAlmostEqual(tt.args.s1, tt.args.s2)
			require.Equal(t, tt.want, got)
		})
	}
}
