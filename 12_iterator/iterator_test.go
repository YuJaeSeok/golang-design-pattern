package iterator

import "testing"

func ExampleIterator() {
	var aggregate Aggregate
	aggregate = NewNumbers(1, 10)

	PrintIterator(aggregate.Iterator())
}

func TestIterator(t *testing.T) {
	ExampleIterator()
}
