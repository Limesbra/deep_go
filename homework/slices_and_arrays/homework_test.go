package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values      []int
	insertIndex int
	deleteIndex int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values:      make([]int, size),
		insertIndex: -1,
		deleteIndex: -1,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	} else {
		q.insertIndex = (q.insertIndex + 1) % len(q.values)
		q.values[q.insertIndex] = value
	}
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	} else {
		q.deleteIndex = (q.deleteIndex + 1) % len(q.values)
		if q.values[q.deleteIndex] != 0 {
			q.values[q.deleteIndex] = 0
		}
	}
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[(q.deleteIndex+1)%len(q.values)]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.insertIndex]
}

func (q *CircularQueue) Empty() bool {
	if q.insertIndex == -1 {
		return true
	}
	if q.deleteIndex == q.insertIndex && q.values[q.insertIndex] == 0 {
		return true
	}
	return q.insertIndex == -1
}

func (q *CircularQueue) Full() bool {
	if q.deleteIndex == -1 && q.insertIndex == len(q.values)-1 {
		return true
	}
	return q.deleteIndex == (q.insertIndex+1)%len(q.values) && q.values[q.deleteIndex] != 0
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
