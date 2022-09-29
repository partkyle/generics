package genmock

type Mock1ArgCall[T any] struct {
	Arg0 T
}

func WasCalledWith1Arg[T comparable](calls []Mock1ArgCall[T], arg0 T) bool {
	for _, i := range calls {
		if i.Arg0 == arg0 {
			return true
		}
	}

	return false
}

type Mock1Arg[T any] struct {
	Calls []Mock1ArgCall[T]
	F     func(T)
}

func (m *Mock1Arg[T]) Call(arg0 T) {
	m.Calls = append(m.Calls, Mock1ArgCall[T]{Arg0: arg0})
	m.F(arg0)
}

type Mock1Arg1Result[T any, R any] struct {
	Calls []Mock1ArgCall[T]
	F     func(T) R
}

func (m *Mock1Arg1Result[T, R]) Call(arg0 T) R {
	m.Calls = append(m.Calls, Mock1ArgCall[T]{Arg0: arg0})
	return m.F(arg0)
}

type Mock1Arg2Results[T any, R any, S any] struct {
	Calls []Mock1ArgCall[T]
	F     func(T) (R, S)
}

func (m *Mock1Arg2Results[T, R, S]) Call(arg0 T) (R, S) {
	m.Calls = append(m.Calls, Mock1ArgCall[T]{Arg0: arg0})
	return m.F(arg0)
}
