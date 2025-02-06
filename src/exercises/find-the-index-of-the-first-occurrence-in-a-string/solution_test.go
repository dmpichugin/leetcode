package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strStr(t *testing.T) {
	t.Parallel()

	const notFound = -1

	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: notFound,
		},
		{
			name: "both empty",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: notFound,
		},
		{
			name: "haystack empty",
			args: args{
				haystack: "",
				needle:   "vanko",
			},
			want: notFound,
		},
		{
			name: "needle empty",
			args: args{
				haystack: "vanko",
				needle:   "",
			},
			want: notFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := strStr(tt.args.haystack, tt.args.needle)
			require.Equal(t, tt.want, got)
		})
	}
}
