package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isUgly(t *testing.T) {
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
				n: 6,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				n: 14,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isUgly(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
