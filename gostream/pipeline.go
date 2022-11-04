package gostream

// type Pipeline[T any] interface {
// 	init(Pipeline[T])
// 	evaluate()
// }

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

func (h *Pipeline[T]) Map(lam func(T) T) Stream[T] {
	return nil
}
