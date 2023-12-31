package factorymethod

import "testing"

func computing(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)

	factory = PlusOperatorFactory{}
	if computing(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}

	factory = MinusOperatorFactory{}
	if computing(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
