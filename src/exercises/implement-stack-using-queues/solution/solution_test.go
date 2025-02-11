package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Parallel()

	stack := Constructor()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for i := 9; i >= 0; i-- {
		val := stack.Pop()
		require.Equal(t, i, val)
	}

	require.True(t, stack.Empty())

	stack.Push(1)
	require.Equal(t, 1, stack.Top())
	require.False(t, stack.Empty())
	require.Equal(t, 1, stack.Pop())
	require.True(t, stack.Empty())
}
