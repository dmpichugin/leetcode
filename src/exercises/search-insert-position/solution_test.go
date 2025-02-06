package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_searchInsert(t *testing.T) {
	t.Parallel()

	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found in middle",
			args: args{
				nums:   deps.GetIntSlice(1, 3, 5, 6),
				target: 5,
			},
			want: 2,
		},
		{
			name: "not found in middle",
			args: args{
				nums:   deps.GetIntSlice(1, 3, 5, 6),
				target: 2,
			},
			want: 1,
		},
		{
			name: "not found in middle",
			args: args{
				nums:   deps.GetIntSlice(1, 3, 5, 6),
				target: 7,
			},
			want: 4,
		},
		{
			name: "not found in end",
			args: args{
				nums:   deps.GetIntSlice(1, 3, 5, 6),
				target: 10,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := searchInsert(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
