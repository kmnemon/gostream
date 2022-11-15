package gostream

import "golang.org/x/exp/constraints"

type sink[T constraints.Ordered] interface {
	begin(int)
	accept(T)
	end()
	cancellationRequested() bool

	setDownStreamSink(sink[T])
}
