
## Getting Started
This is a library for go that use like stream style programming  
go get github.com/kmnemon/gostream  

## Features
- Support Slice and Map
  
- Map(func(T) T) stream[T]
- Reduce(func(T, T) T) stream[T]
- ReduceWithInitValue(T, func(T, T) T) stream[T]
- ForEach(func(T))
- Sorted() stream[T]
- SortedWith(func(T, T) bool) stream[T]
- Filter(func(T) bool) stream[T]
- Limit(int) stream[T]
- FindFirst() T
- ToList() []T
-	DistinctWith(func(T, T) bool) stream[T]

## Example
```
  a := []int{6, 5, 3, 4, 5}

  x := StreamOf(a).  
    ReduceWithInitValue(1, func(x int, y int) int {  
      return x + y  
    }).  
    ToList()  

  fmt.Println(x)
```
-----------
[7 12 15 19 24]

```
ages := map[string]int{
		"r": 1000,
		"v": 2000,
		"c": 500,
	}

	x := StreamOfMap(ages).Map(func(e EntrySet[string, int]) EntrySet[string, int] {
		e.V = e.V + 1
		return e
	}).ToList()
```
-----------
{"r": 1001},
{"v": 2001},
{"c": 501}

## License

This software is released under the GPL-3.0 license.