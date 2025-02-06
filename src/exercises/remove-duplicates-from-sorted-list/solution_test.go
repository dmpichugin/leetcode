package solution

import (
	"testing"

	. "github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_deleteDuplicates(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "with duplicates in beginning",
			args: args{
				head: GetIntList(1, 1, 2, 3),
			},
			want: GetIntList(1, 2, 3),
		},
		{
			name: "with duplicates in beginning and ending",
			args: args{
				head: GetIntList(1, 1, 2, 3, 3),
			},
			want: GetIntList(1, 2, 3),
		},
		{
			name: "empty",
			args: args{
				head: GetIntList(),
			},
			want: GetIntList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := deleteDuplicates(tt.args.head)
			require.True(t, tt.want.Equals(got))
		})
	}
}
