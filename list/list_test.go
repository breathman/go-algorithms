package list

import (
	"gotest.tools/assert"
	"testing"
)

func TestList_Size(t *testing.T) {
	l := NewList(1, 2, 3)
	l.AddHead(4)
	assert.Equal(t, 4, l.Size())
}

func TestList_IsEmpty(t *testing.T) {
	l := NewList()
	assert.Check(t, l.IsEmpty())
	l.AddHead(2)
	assert.Check(t, !l.IsEmpty())
}

func TestList_AddHead(t *testing.T) {
	list := NewList(2, 3)

	list.AddHead(1)
	assert.Equal(t, list.Size(), 3)

	list.AddHead(2)
	assert.Equal(t, list.Size(), 4)
	assert.Equal(t, list.head.value, 2)
}

func TestList_AddTail(t *testing.T) {
	l := NewList(2, 3)

	l.AddTail(1)
	assert.Equal(t, l.head.value, 2)
	assert.Equal(t, l.Size(), 3)
}

func TestList_Print(t *testing.T) {
	l := NewList(4, 1, 2)
	l.Print()
}

func TestNewList(t *testing.T) {
	l := NewList(1, 2, 3, 4, 5)
	assert.Equal(t, 5, l.Size())
	assert.Equal(t, 1, l.head.value)
}
