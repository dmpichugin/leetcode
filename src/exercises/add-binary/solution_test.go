package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addBinary(t *testing.T) {
	t.Parallel()

	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "without overflow",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "with overflow",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := addBinary(tt.args.a, tt.args.b)
			require.Equal(t, tt.want, got)
		})
	}
}
