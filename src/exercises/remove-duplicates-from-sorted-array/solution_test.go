package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantNumber int
		wantSlice  []int
	}{
		{
			name: "without duplicates",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3),
			},
			wantNumber: 3,
			wantSlice:  deps.GetIntSlice(1, 2, 3),
		},
		{
			name: "with 1 duplicate in end",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3, 3),
			},
			wantNumber: 3,
			wantSlice:  deps.GetIntSlice(1, 2, 3),
		},
		{
			name: "with 1 duplicate in begin",
			args: args{
				nums: deps.GetIntSlice(1, 1, 2, 3),
			},
			wantNumber: 3,
			wantSlice:  deps.GetIntSlice(1, 2, 3),
		},
		{
			name: "all duplicates",
			args: args{
				nums: deps.GetIntSlice(1, 1, 1, 1, 1),
			},
			wantNumber: 1,
			wantSlice:  deps.GetIntSlice(1),
		},
		{
			name: "empty",
			args: args{
				nums: deps.GetIntSlice(),
			},
			wantNumber: 0,
			wantSlice:  deps.GetIntSlice(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := removeDuplicates(tt.args.nums)

			require.Equal(t, tt.wantNumber, got)
			require.Equal(t, tt.wantSlice, tt.args.nums[:got])
		})
	}
}
