package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_plusOne(t *testing.T) {
	t.Parallel()

	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				digits: deps.GetIntSlice(1, 2, 3),
			},
			want: deps.GetIntSlice(1, 2, 4),
		},
		{
			name: "case 2",
			args: args{
				digits: deps.GetIntSlice(4, 3, 2, 1),
			},
			want: deps.GetIntSlice(4, 3, 2, 2),
		},
		{
			name: "overflow",
			args: args{
				digits: deps.GetIntSlice(9),
			},
			want: deps.GetIntSlice(1, 0),
		},
		{
			name: "zero",
			args: args{
				digits: deps.GetIntSlice(0),
			},
			want: deps.GetIntSlice(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := plusOne(tt.args.digits)
			require.Equal(t, tt.want, got)
		})
	}
}
