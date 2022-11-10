package gostream

type StateType int

const (
	Head StateType = iota
	StatelessOp
	StatefulOp
)

type Pipeline[T any, R any] struct {
	PreviousStage *Pipeline[T, R]
	NextStage     *Pipeline[T, R]
	Depth         int
	StreamOpFlag  StateType
	StreamSink    Sink[T]
}

func (p *Pipeline[T, R]) New(previousStage *Pipeline[T, R], opFlag StateType) {
	if opFlag == Head {
		p.PreviousStage = nil
		p.Depth = 0
	} else {
		p.PreviousStage.NextStage = p
		p.PreviousStage = previousStage
		p.Depth = p.PreviousStage.Depth + 1
	}

	p.StreamOpFlag = opFlag

}

func (p *Pipeline[T, R]) OpWrapSink(downstream Sink[T]) Sink[T] {

	return nil
}

func (p *Pipeline[T, R]) Map(mapper func(T) R) Stream[T, R] {
	statelessPipe := Pipeline[T, R]{}
	statelessPipe.New(p, StatelessOp)

	return &statelessPipe
}

func (p *Pipeline[T, R]) Sink(downStream Sink[T]) Sink[T] {
	return nil
}
