package main

import "fmt"

func main() {
	m1 := NewMatrix(5, 5)
	m2 := NewMatrix(6, 9)

	prettyPrint(m1)
	prettyPrint(m2)
}

func prettyPrint(m [][]float64) {
	fmt.Println("------------------------")
	for _, r := range m {
		fmt.Println(r)
	}
}
