package solution

import (
	"testing"

	. "github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_mergeTwoLists(t *testing.T) {
	t.Parallel()

	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			require.True(t, tt.want.Equals(got))
		})
	}
}
