package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_moveZeroes(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3),
			},
			want: deps.GetIntSlice(1, 2, 3),
		},
		{
			name: "case 2",
			args: args{
				nums: deps.GetIntSlice(0, 1, 0),
			},
			want: deps.GetIntSlice(1, 0, 0),
		},
		{
			name: "case 3",
			args: args{
				nums: deps.GetIntSlice(0, 0, 0, 0, 1),
			},
			want: deps.GetIntSlice(1, 0, 0, 0, 0),
		},
		{
			name: "case 4",
			args: args{
				nums: deps.GetIntSlice(1, 0, 2, 0, 4, 5),
			},
			want: deps.GetIntSlice(1, 2, 4, 5, 0, 0),
		},
		{
			name: "case 5",
			args: args{
				nums: deps.GetIntSlice(0),
			},
			want: deps.GetIntSlice(0),
		},
		{
			name: "case 6",
			args: args{
				nums: deps.GetIntSlice(1),
			},
			want: deps.GetIntSlice(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			moveZeroes(tt.args.nums)
			require.Equal(t, tt.want, tt.args.nums)
		})
	}
}
