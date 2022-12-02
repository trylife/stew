package v1_test

import (
	v1 "github.com/trylife/stew/cs/dsa/ds/stack/v1"
	"testing"
)

func TestStack(t *testing.T) {
	s := v1.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s)
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}
