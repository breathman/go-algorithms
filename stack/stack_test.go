package stack

import (
	"gotest.tools/assert"
	"testing"
)

func TestStack(t *testing.T) {

	st := New()
	assert.Assert(t, st.IsEmpty())

	st.Push(1)
	st.Push(5)
	st.Push(3)

	assert.Assert(t, !st.IsEmpty())
	assert.Equal(t, 3, st.Size())

	item, err := st.Pop()
	assert.Equal(t, 3, item)
	assert.NilError(t, err)
	assert.Equal(t, 2, st.Size())

	top, err := st.Top()
	assert.Equal(t, 5, top)
	assert.NilError(t, err)
}
