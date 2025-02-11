package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfFour(t *testing.T) {
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
			name: "4",
			args: args{
				n: 4,
			},
			want: true,
		},
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "16",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "15",
			args: args{
				n: 15,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isPowerOfFour(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
