package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hammingWeight(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 11,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				n: 128,
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				n: 2147483645,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := hammingWeight(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
