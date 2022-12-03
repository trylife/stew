package v2

// Stack golang stack
type Stack []int

func NewStack() Stack {
	s := make(Stack, 0)
	return s
}

func Push(s Stack, v int) Stack {
	return append(s, v)
}

func Pop(s Stack) (Stack, int, bool) {
	l := len(s)
	if l == 0 {
		return s, 0, false
	}

	re := s[l-1]
	s = s[:l-1]

	return s, re, true
}
