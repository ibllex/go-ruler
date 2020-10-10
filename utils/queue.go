package utils

// Queue FIFO (First In First Out) data struct
type Queue struct {
	values []interface{}
}

// Push add item to queue's end
func (q *Queue) Push(value interface{}) {
	q.values = append(q.values, value)
}

// Pop fetch the queue header element and return
func (q *Queue) Pop() interface{} {
	v := q.values[0]
	q.values = q.values[1:]
	return v
}

// IsEmpty is queue empty
func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}

// Size return the queue's elements count
func (q *Queue) Size() int {
	return len(q.values)
}
