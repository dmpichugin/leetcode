package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isHappy(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				n: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isHappy(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
