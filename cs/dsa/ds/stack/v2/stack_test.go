package v2_test

import (
	v2 "github.com/trylife/stew/cs/dsa/ds/stack/v2"
	"testing"
)

func TestStack(t *testing.T) {
	s := v2.Stack{}
	s = v2.Push(s, 1)
	s = v2.Push(s, 2)
	s = v2.Push(s, 3)
	t.Log(s)

	t.Log(v2.Pop(s))
	t.Log(v2.Pop(s))
	t.Log(v2.Pop(s))
	t.Log(v2.Pop(s))

	s, v, ok := v2.Pop(s)
	t.Log(s, v, ok)

	s, v, ok = v2.Pop(s)
	t.Log(s, v, ok)

	s, v, ok = v2.Pop(s)
	t.Log(s, v, ok)

	s, v, ok = v2.Pop(s)
	t.Log(s, v, ok)
}
