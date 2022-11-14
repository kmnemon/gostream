package gostream

type stateType int

const (
	head stateType = iota
	statelessOp
	statefulOp
	terminalOp
)

type abstractPipeline[T any] interface {
	wrapSink(sink[T]) sink[T]
	copyInto(sink[T], []T)
}

type pipeline[T any] struct {
	previousStage *pipeline[T]
	nextStage     *pipeline[T]
	depth         int
	streamOpFlag  stateType
	streamSink    sink[T]
}

func (p *pipeline[T]) new(previousStage *pipeline[T], opFlag stateType, sink sink[T]) {
	if opFlag == head {
		p.previousStage = nil
		p.depth = 0
	} else {
		p.previousStage.nextStage = p
		p.previousStage = previousStage
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
		nil,
		mapper,
	}
	statelessPipe.new(p, statelessOp, &s)

	return &statelessPipe
}

func (p *pipeline[T]) ForEach(mapper func(T)) {
	terminalPipe := pipeline[T]{}
	s := forEachSink[T]{
		mapper,
	}
	terminalPipe.new(p, terminalOp, &s)

	sink := p.wrapSink(&s)
	p.copyInto(sink, nil)
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
