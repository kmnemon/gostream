package gostream

import "golang.org/x/exp/constraints"

type stateType int

const (
	head stateType = iota
	statelessOp
	statefulOp
)

type abstractPipeline[T constraints.Ordered] interface {
	wrapSink(sink[T]) sink[T]
	copyInto(sink[T], []T)
	evaluate()
}

type pipeline[T constraints.Ordered] struct {
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
		nil,
		mapper,
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
	// statefulPipe := pipeline[T]{}
	// s := sortingSink[T]{}
	return nil
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
