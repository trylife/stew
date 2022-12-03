package alg3_1_test

import (
	"fmt"
	v2 "github.com/trylife/stew/cs/dsa/ds/stack/v2"
	"testing"
)

func conversion(decimal int) {

}

// DS_3_2_1 数制转换
func TestConversion(t *testing.T) {
	decimal := 1348
	s := v2.NewStack()
	s = v2.Push(s, decimal)
	var ok bool

	for len(s) > 0 {
		s, decimal, ok = v2.Pop(s)

		n := decimal / 8
		m := decimal % 8
		fmt.Println(decimal, n, m, ok)
		if n > 0 {
			s = v2.Push(s, n)
		}
	}

}
