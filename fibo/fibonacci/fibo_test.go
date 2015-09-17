package fibonacci

import (
	"strconv"
	"testing"
)

func TestFib(t *testing.T) {
	result := "01123"
	var fib string
	for i := 0; i < 5; i++ {
		fib = fib + strconv.Itoa(Fib(i))
	}
	if fib != result {
		t.Error("Expected", result, " got ", fib)
	}
}
