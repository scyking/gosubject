// package name + _test 属于黑盒测试，仅能测试公开方法
package algorithms_test

import (
	"fmt"
	"gosubject/algorithms"
	"testing"
)

func TestRecursion(t *testing.T) {
	println("recursion test start:")
	n := 10
	for n >= 0 {
		fmt.Printf("n:%d,f(n):%d \n", n, algorithms.Recursion(n))
		n--
	}
}

func TestLoop(t *testing.T) {
	println("loop test start:")
	n := 10
	for n >= 0 {
		fmt.Printf("n:%d,f(n):%d \n", n, algorithms.Loop(n))
		n--
	}
}
