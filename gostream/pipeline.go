package gostream

type stateType int

const (
	head stateType = iota
	statelessOp
	statefulOp
)

type abstractPipeline[T any] interface {
	wrapSink(sink[T]) sink[T]
	copyInto(sink[T], []T)
	evaluate()
}

type pipeline[T any] struct {
	previousStage *pipeline[T]
	nextStage     *pipeline[T]
	sourceStage   *pipeline[T]
	depth         int
	streamOpFlag  stateType
	streamSink    sink[T]
	sourceData    []T
}

func (p *pipeline[T]) init(previousStage *pipeline[T], opFlag stateType, sink sink[T]) {
	if opFlag == head {
		p.previousStage = nil
		p.sourceStage = p
		p.depth = 0
	} else {
		p.previousStage = previousStage
		p.previousStage.nextStage = p
		p.sourceStage = previousStage.sourceStage
		p.depth = p.previousStage.depth + 1
		p.streamSink = sink
	}

	p.streamOpFlag = opFlag

}

func (p *pipeline[T]) opWrapSink(downstream sink[T]) sink[T] {
	p.streamSink.setDownStreamSink(downstream)
	return p.streamSink
}

func (p *pipeline[T]) Map(mapper func(T) T) stream[T] {
	statelessPipe := pipeline[T]{}
	s := mapSink[T]{
		mapper,
		nil,
	}
	statelessPipe.init(p, statelessOp, &s)

	return &statelessPipe
}

func (p *pipeline[T]) Reduce(binaryOperator func(T, T) T) stream[T] {
	statelessPipe := pipeline[T]{}
	s := reduceSink[T]{
		binaryOperator: binaryOperator,
		downstream:     nil,
		isFirstValue:   true,
	}
	statelessPipe.init(p, statelessOp, &s)

	return &statelessPipe
}

func (p *pipeline[T]) ReduceWithInitValue(i T, binaryOperator func(T, T) T) stream[T] {
	statelessPipe := pipeline[T]{}
	s := reduceSink[T]{
		binaryOperator: binaryOperator,
		downstream:     nil,
		i:              i,
		isFirstValue:   false,
	}
	statelessPipe.init(p, statelessOp, &s)

	return &statelessPipe
}

func (p *pipeline[T]) ForEach(mapper func(T)) {
	s := forEachSink[T]{
		mapper,
	}

	p.evaluate(&s)

}

func (p *pipeline[T]) Sorted() stream[T] {
	statefulPipe := pipeline[T]{}
	s := sortingSink[T]{}
	statefulPipe.init(p, statefulOp, &s)

	return &statefulPipe
}

func (p *pipeline[T]) SortedWith(less func(T, T) bool) stream[T] {
	statefulPipe := pipeline[T]{}
	s := sortingSink[T]{
		less,
		nil,
		nil,
	}
	statefulPipe.init(p, statefulOp, &s)

	return &statefulPipe
}

func (p *pipeline[T]) Filter(predicate func(T) bool) stream[T] {
	statelessPipe := pipeline[T]{}
	s := filterSink[T]{
		predicate,
		nil,
	}
	statelessPipe.init(p, statefulOp, &s)

	return &statelessPipe
}

func (p *pipeline[T]) Limit(maxSize int) stream[T] {
	statelessPipe := pipeline[T]{}
	s := limitSink[T]{
		maxSize,
		nil,
	}
	statelessPipe.init(p, statefulOp, &s)

	return &statelessPipe
}

func (p *pipeline[T]) FindFirst() T {
	s := findFirstSink[T]{}
	p.evaluate(&s)
	return s.result
}

func (p *pipeline[T]) ToList() []T {
	s := toListSink[T]{}
	p.evaluate(&s)
	return s.list
}

func (p *pipeline[T]) evaluate(s sink[T]) {
	p.copyInto(p.wrapSink(s), p.sourceStage.sourceData)
}

func (p *pipeline[T]) wrapSink(sink sink[T]) sink[T] {
	for ; p.depth > 0; p = p.previousStage {
		sink = p.opWrapSink(sink)
	}

	return sink
}

func (p *pipeline[T]) copyInto(wrapSink sink[T], slice []T) {
	wrapSink.begin(len(slice))
	for _, s := range slice {
		wrapSink.accept(s)
	}
	wrapSink.end()
}
