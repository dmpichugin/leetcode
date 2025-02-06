package solution

import (
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_longestCommonPrefix(t *testing.T) {
	t.Parallel()

	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple case",
			args: args{
				strs: deps.GetStringSlice("flower", "flow", "flight"),
			},
			want: "fl",
		},
		{
			name: "empty",
			args: args{
				strs: deps.GetStringSlice("", "", ""),
			},
			want: "",
		},
		{
			name: "one word",
			args: args{
				strs: deps.GetStringSlice("flower"),
			},
			want: "flower",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := longestCommonPrefix(tt.args.strs)
			require.Equal(t, tt.want, got)
		})
	}
}
