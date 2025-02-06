package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLastWord(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				s: "   fly me   to   the moon  ",
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				s: "luffy is still joyboy",
			},
			want: 6,
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := lengthOfLastWord(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
