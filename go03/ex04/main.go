package main

import "fmt"

func main() {
	i1 := 42
	fmt.Println("Before changes:", i1)
	doublePtr(&i1)
	fmt.Println("After first call:", i1)
	i1 = doubleRet(i1)
	fmt.Println("After second change", i1)
}

func doublePtr(n *int) {
	*n *= 2
}

func doubleRet(n int) int {
	return n * 2
}
