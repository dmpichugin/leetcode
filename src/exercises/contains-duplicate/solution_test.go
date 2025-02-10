package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_containsDuplicate(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3),
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				nums: deps.GetIntSlice(1, 2, 3, 3),
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				nums: deps.GetIntSlice(1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := containsDuplicate(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
