package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_countBits(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				n: 2,
			},
			want: deps.GetIntSlice(0, 1, 1),
		},
		{
			name: "case 2",
			args: args{
				n: 5,
			},
			want: deps.GetIntSlice(0, 1, 1, 2, 1, 2),
		},
		{
			name: "case 3",
			args: args{
				n: 0,
			},
			want: deps.GetIntSlice(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := countBits(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
