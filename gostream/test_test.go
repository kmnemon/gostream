package gostream

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	StreamOf(a).
		Map(func(x int) int {
			return x * 3
		}).
		ForEach(func(x int) {
			fmt.Println(x)
		})
}
