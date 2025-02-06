package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_singleNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				nums: deps.GetIntSlice(1, 1, 3),
			},
			want: 3,
		},
		{
			name: "without",
			args: args{
				nums: deps.GetIntSlice(1, 3, 3, 1),
			},
			want: 0,
		},
		{
			name: "empty",
			args: args{
				nums: deps.GetIntSlice(),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := singleNumber(tt.args.nums)
			require.Equal(t, got, tt.want)
		})
	}
}
