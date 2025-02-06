package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValid(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string ",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "valid case",
			args: args{
				s: "[{}]",
			},
			want: true,
		},
		{
			name: "invalid order",
			args: args{
				s: "[{]}",
			},
			want: false,
		},
		{
			name: "miss element",
			args: args{
				s: "[{}",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isValid(tt.args.s)
			require.Equal(t, got, tt.want)
		})
	}
}
