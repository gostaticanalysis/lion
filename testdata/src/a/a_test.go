package a_test

import (
	"a"
	"testing"
)

func Test_F1(t *testing.T) {
	a.F1()
}

func Test_T_M1(t *testing.T) {
	(a.T{}).M1()
}
