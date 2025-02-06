package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantSlice []int
	}{
		{
			name: "without value",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3),
				val:  5,
			},
			want:      3,
			wantSlice: deps.GetIntSlice(1, 2, 3),
		},
		{
			name: "with value in begin",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3),
				val:  1,
			},
			want:      2,
			wantSlice: deps.GetIntSlice(2, 3),
		},
		{
			name: "with duplicate value in begin",
			args: args{
				nums: deps.GetIntSlice(1, 1, 1, 2, 3),
				val:  1,
			},
			want:      2,
			wantSlice: deps.GetIntSlice(2, 3),
		},
		{
			name: "with duplicate value in end",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3, 3, 3, 3),
				val:  3,
			},
			want:      2,
			wantSlice: deps.GetIntSlice(1, 2),
		},
		{
			name: "with duplicate value in middle",
			args: args{
				nums: deps.GetIntSlice(1, 2, 2, 2, 3, 3, 3, 3),
				val:  2,
			},
			want:      5,
			wantSlice: deps.GetIntSlice(1, 3, 3, 3, 3),
		},
		{
			name: "empty nums",
			args: args{
				nums: deps.GetIntSlice(),
				val:  10,
			},
			want:      0,
			wantSlice: deps.GetIntSlice(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := removeElement(tt.args.nums, tt.args.val)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantSlice, tt.args.nums[:got])
		})
	}
}
