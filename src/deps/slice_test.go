package deps

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetIntSlice(t *testing.T) {
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
			name: "empty",
			args: args{
				nums: []int{},
			},
			want: nil,
		},
		{
			name: "single",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "multiple",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := GetIntSlice(tt.args.nums...)
			require.Equal(t, tt.want, got)
		})
	}
}
