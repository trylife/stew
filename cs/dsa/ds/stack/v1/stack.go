package v1

// Stack golang stack
type Stack []int

func NewStack() Stack {
	s := make(Stack, 0)
	return s
}

func (s Stack) Push(i int) {
	s = append(s, i)
}

func (s Stack) Pop() (re int, ok bool) {
	l := s.Len()
	if l == 0 {
		return 0, false
	}

	re = s[l-1]
	s = s[:l-1]

	return re, true
}

func (s Stack) Len() int {
	return len(s)
}
