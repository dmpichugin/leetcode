package deps

import (
	"sort"
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

func TestSumSlice(t *testing.T) {
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
			name: "case 1",
			args: args{
				nums: GetIntSlice(),
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				nums: GetIntSlice(1, 2, 3),
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				nums: GetIntSlice(1),
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				nums: GetIntSlice(-2, -1, 0, 1, 2),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := SumSlice(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGetDigits(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				num: 10,
			},
			want: GetIntSlice(1, 0),
		},
		{
			name: "case 2",
			args: args{
				num: 0,
			},
			want: GetIntSlice(0),
		},
		{
			name: "case 3",
			args: args{
				num: 102,
			},
			want: GetIntSlice(1, 0, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := GetDigits(tt.args.num)

			sort.Ints(tt.want)
			sort.Ints(got)

			require.Equal(t, tt.want, got)
		})
	}
}
