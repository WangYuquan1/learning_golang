package hello_test

import (
	"testing"

	"learning_golang/app/hello"
)

func TestSum(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := hello.Sum(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}

}

func Sum(a, b int) {
	panic("unimplemented")
}
