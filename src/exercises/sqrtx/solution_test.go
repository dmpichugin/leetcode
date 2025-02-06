package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mySqrt(t *testing.T) {
	t.Parallel()

	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal case",
			args: args{
				s: 4,
			},
			want: 2,
		},
		{
			name: "zero",
			args: args{
				s: 0,
			},
			want: 0,
		},
		{
			name: "not exactly",
			args: args{
				s: 8,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := mySqrt(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
