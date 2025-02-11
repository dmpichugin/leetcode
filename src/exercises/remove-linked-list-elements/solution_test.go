package solution

import (
	"testing"

	. "github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_removeElements(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "with 3 vals elements",
			args: args{
				head: GetIntList(5, 1, 2, 5, 3, 4, 5),
				val:  5,
			},
			want: GetList(1, 2, 3, 4),
		},
		{
			name: "with 0 vals elements",
			args: args{
				head: GetIntList(1, 2, 3, 4, 6),
				val:  5,
			},
			want: GetList(1, 2, 3, 4, 6),
		},
		{
			name: "empty list",
			args: args{
				head: GetIntList(),
				val:  5,
			},
			want: GetList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := removeElements(tt.args.head, tt.args.val)
			require.Truef(t, tt.want.Equals(got), "\n want: %s\n got:  %s\n", tt.want.String(), got.String())
		})
	}
}
