package utils

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := &Queue{}
	values := []interface{}{
		0, 1, "hello, world", 1.0, nil,
	}

	for _, v := range values {
		q.Push(v)
	}

	if q.Size() != len(values) {
		t.Fatalf("queue size error, expected %d, but %d got.", len(values), q.Size())
	}

	for _, v := range values {
		got := q.Pop()
		if v != got {
			t.Fatalf("queue error, expected %v, but %v got.", v, got)
		}
	}

	if q.Size() != 0 {
		t.Fatalf("queue size error, expected 0, but %d got.", q.Size())
	}

	if q.IsEmpty() != true {
		t.Fatalf("queue is not empty.")
	}
}
