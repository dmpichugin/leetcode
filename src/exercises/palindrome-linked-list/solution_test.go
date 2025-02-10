package solution

import (
	"testing"

	. "github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
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
			name: "case 1 palindrome",
			args: args{
				head: GetList(1, 2, 2, 1),
			},
			want: true,
		},
		{
			name: "case 2 palindrome",
			args: args{
				head: GetList(0, 0),
			},
			want: true,
		},
		{
			name: "case 3 not palindrome",
			args: args{
				head: GetList(0, 0, 1),
			},
			want: false,
		},
		{
			name: "case 4 palindrome",
			args: args{
				head: GetList(1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isPalindrome(tt.args.head)
			require.Equal(t, tt.want, got)
		})
	}
}
