package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums1: deps.GetIntSlice(1, 2, 3, 0, 0, 0),
				m:     3,
				nums2: deps.GetIntSlice(2, 5, 6),
				n:     3,
			},
			want: deps.GetIntSlice(1, 2, 2, 3, 5, 6),
		},
		{
			name: "case 1",
			args: args{
				nums1: deps.GetIntSlice(1),
				m:     1,
				nums2: deps.GetIntSlice(),
				n:     0,
			},
			want: deps.GetIntSlice(1),
		},
		{
			name: "case 1",
			args: args{
				nums1: deps.GetIntSlice(0),
				m:     0,
				nums2: deps.GetIntSlice(1),
				n:     1,
			},
			want: deps.GetIntSlice(1),
		},
		{
			name: "both empty",
			args: args{
				nums1: deps.GetIntSlice(),
				m:     0,
				nums2: deps.GetIntSlice(0),
				n:     0,
			},
			want: deps.GetIntSlice(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
