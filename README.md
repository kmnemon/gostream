
## Getting Started
This is a library for go that use like stream style programming  
go get github.com/kmnemon/gostream  

## Features
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

## License

This software is released under the GPL-3.0 license.