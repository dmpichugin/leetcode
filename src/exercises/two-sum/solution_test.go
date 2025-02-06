package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums:   deps.GetIntSlice(2, 7, 11, 15),
				target: 9,
			},
			want: deps.GetIntSlice(0, 1),
		},
		{
			name: "case 2",
			args: args{
				nums:   deps.GetIntSlice(3, 2, 4),
				target: 6,
			},
			want: deps.GetIntSlice(1, 2),
		},
		{
			name: "case 3",
			args: args{
				nums:   deps.GetIntSlice(3, 3),
				target: 6,
			},
			want: deps.GetIntSlice(0, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := twoSum(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
