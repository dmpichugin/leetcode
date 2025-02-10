package solution

import (
	"math"
	"testing"

	"github.com/dmpichugin/leetcode/src/deps"
	"github.com/stretchr/testify/require"
)

func Test_reverseBits(t *testing.T) {
	t.Parallel()

	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "case 1",
			args: args{
				num: 3,
			},
			want: 3221225472,
		},
		{
			name: "case 2",
			args: args{
				num: 5,
			},
			want: 2684354560,
		},
		{
			name: "case 3",
			args: args{
				num: 8,
			},
			want: 268435456,
		},
		{
			name: "max uint32",
			args: args{
				num: math.MaxUint32,
			},
			want: math.MaxUint32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := reverseBits(tt.args.num)
			require.Equalf(t, tt.want, got, "want %s got %s", deps.FormatBinary32(tt.want), deps.FormatBinary32(got))
		})
	}
}
