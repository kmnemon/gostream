package gostream

import (
	"testing"
)

func TestMap(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Map(func(x int) int {
			return x + 1
		}).
		ToList()

	expect := []int{7, 6, 4, 5, 6}

	if !equalSliceHelper(expect, x) {
		t.Error("Map operator has some problem")
	}
}

func TestMapWithMap(t *testing.T) {
	ages := map[string]int{
		"r": 1000,
		"v": 2000,
		"c": 500,
	}

	x := StreamOfMap(ages).Map(func(e EntrySet[string, int]) EntrySet[string, int] {
		e.V = e.V + 1
		return e
	}).ToList()

	expect := []EntrySet[string, int]{
		{K: "r", V: 1001},
		{K: "v", V: 2001},
		{K: "c", V: 501},
	}

	if !equalSliceHelper(expect, x) {
		t.Error("MapWithMap operator has some problem")
	}
}

func TestDistinct(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Distinct().
		ToList()

	expect := []int{6, 5, 3, 4}

	if !equalSliceHelper(expect, x) {
		t.Error("Distinct operator has some problem")
	}
}

func TestReduce(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Reduce(func(x int, y int) int {
			return x + y
		}).
		ToList()

	expect := []int{6, 11, 14, 18, 23}
	if !equalSliceHelper(expect, x) {
		t.Error("Reduce operator has some problem")
	}
}

func TestReduceWithInitValue(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		ReduceWithInitValue(1, func(x int, y int) int {
			return x + y
		}).
		ToList()

	expect := []int{7, 12, 15, 19, 24}
	if !equalSliceHelper(expect, x) {
		t.Error("ReduceWithInitValue operator has some problem")
	}
}

func TestSorted(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Sorted().
		ToList()

	expect := []int{3, 4, 5, 5, 6}

	if !equalSliceHelper(expect, x) {
		t.Error("Sorted operator has some problem")
	}
}

func TestFilter(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Filter(func(a int) bool {
			return a > 4
		}).
		ToList()

	expect := []int{6, 5, 5}

	if !equalSliceHelper(expect, x) {
		t.Error("Filter operator has some problem")
	}
}

func TestLimit(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Limit(3).
		ToList()

	expect := []int{6, 5, 3}

	if !equalSliceHelper(expect, x) {
		t.Error("Limit operator has some problem")
	}
}

func TestFindFirst(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		FindFirst()

	if !(x == 6) {
		t.Error("FindFirst operator has some problem")
	}
}

// b := []AA{
// 	{1, 2},
// 	{3, 4},
// }

// less := func(a AA, b AA) bool {
// 	return b.a < a.a
// }

// StreamOf(b).
// 	// Map(func(x int) int {
// 	// 	return x * 3
// 	// }).
// 	SortedWith(less).
// 	ForEach(func(x AA) {
// 		fmt.Println(x)
// 	})

// ages := map[string]int{
// 	"rr": 1000,
// 	"vv": 2000,
// 	"cc": 500,
// }

// StreamOfMap(ages).Map(func(e EntrySet[string, int]) EntrySet[string, int] {
// 	e.V = e.V + 1
// 	return e
// }).ForEach(func(e EntrySet[string, int]) {
// 	fmt.Println(e)
// })

func equalSliceHelper[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
