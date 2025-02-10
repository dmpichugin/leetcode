package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addDigits(t *testing.T) {
	t.Parallel()
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "38",
			args: args{
				num: 38,
			},
			want: 2,
		},
		{
			name: "0",
			args: args{
				num: 0,
			},
			want: 0,
		},
		{
			name: "101",
			args: args{
				num: 101,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := addDigits(tt.args.num)
			require.Equal(t, tt.want, got)
		})
	}
}
