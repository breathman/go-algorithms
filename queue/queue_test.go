package queue

import (
	"gotest.tools/assert"
	"testing"
)

func TestQueue(t *testing.T) {

	q := New()
	assert.Assert(t, q.IsEmpty())

	q.Push(1)
	q.Push(5)
	q.Push(3)

	assert.Assert(t, !q.IsEmpty())
	assert.Equal(t, 3, q.Size())

	item, err := q.Pop()
	assert.Equal(t, 1, item)
	assert.NilError(t, err)
	assert.Equal(t, 2, q.Size())

	top, err := q.Top()
	assert.Equal(t, 5, top)
	assert.NilError(t, err)

}
