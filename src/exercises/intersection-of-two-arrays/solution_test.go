package solution

import (
	"sort"
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_intersection(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "with duplicates have intersection",
			args: args{
				nums1: deps.GetIntSlice(1, 2, 3, 3),
				nums2: deps.GetIntSlice(1, 3),
			},
			want: deps.GetIntSlice(1, 3),
		},
		{
			name: "with duplicates have not intersection",
			args: args{
				nums1: deps.GetIntSlice(1, 2, 3, 3),
				nums2: deps.GetIntSlice(4, 5),
			},
			want: deps.GetIntSlice(),
		},
		{
			name: "without duplicates have intersection",
			args: args{
				nums1: deps.GetIntSlice(1, 2, 3),
				nums2: deps.GetIntSlice(3, 4, 5),
			},
			want: deps.GetIntSlice(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := intersection(tt.args.nums1, tt.args.nums2)

			sort.Ints(got)
			sort.Ints(tt.want)

			require.Equal(t, tt.want, got)
		})
	}
}
