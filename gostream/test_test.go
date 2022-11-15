package gostream

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	StreamOf(a).Map(func(x int) int {
		return x * 2
	}).ForEach(func(x int) {
		fmt.Println(x)
	})
}
