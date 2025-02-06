package solution

import (
	"testing"

	. "github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_hasCycle(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "without cycle",
			args: args{
				head: GetIntList(1, 2, 3),
			},
			want: false,
		},
		{
			name: "without cycle in end",
			args: args{
				head: GetIntList(1, 2, 3).MakeCycle(),
			},
			want: true,
		},
		{
			name: "without cycle in middle",
			args: args{
				head: GetIntList(1, 2, 3, 4, 5).MakeCycleInMiddle(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := hasCycle(tt.args.head)
			require.Equal(t, tt.want, got)
		})
	}
}
