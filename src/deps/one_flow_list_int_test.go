package deps

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListNode_Equals(t *testing.T) {
	t.Parallel()

	type args struct {
		a *ListNode
		b *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equals",
			args: args{
				a: GetIntList(1, 2, 3),
				b: GetIntList(1, 2, 3),
			},
			want: true,
		},
		{
			name: "not equals",
			args: args{
				a: GetIntList(1, 2, 3),
				b: GetIntList(1, 2),
			},
			want: false,
		},
		{
			name: "both nil",
			args: args{
				a: nil,
				b: nil,
			},
			want: true,
		},
		{
			name: "receiver is nil",
			args: args{
				a: nil,
				b: GetIntList(1, 2),
			},
			want: false,
		},
		{
			name: "second is nil",
			args: args{
				a: GetIntList(1, 2),
				b: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, tt.args.a.Equals(tt.args.b))
		})
	}
}
