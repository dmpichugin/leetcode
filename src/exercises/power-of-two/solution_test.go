package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfTwo(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{n: 1},
			want: true,
		},
		{
			name: "two",
			args: args{n: 2},
			want: true,
		},
		{
			name: "four",
			args: args{n: 4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isPowerOfTwo(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
