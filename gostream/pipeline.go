package gostream

type StateType int

const (
	Head StateType = iota
	StatelessOp
	StatefulOp
)

type Pipeline[T any] struct {
	PreviousStage *Pipeline[T]
	NextStage     *Pipeline[T]
	Depth         int
	StreamOpFlag  StateType
	StreamSink    Sink[T]
}

func (p *Pipeline[T]) New(previousStage *Pipeline[T], opFlag StateType) {
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

func (p *Pipeline[T]) Map(lam func(T) T) Stream[T] {
	return nil
}

func (p *Pipeline[T]) Sink(downStream Sink[T]) Sink[T] {
	return nil
}
