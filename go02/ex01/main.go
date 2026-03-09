package main

func main() {
	m1 := NewMatrix(5, 5)
	m2 := NewMatrix(6, 9)

	PrintMatrix(m1)
	PrintMatrix(Transpose(m1))
	PrintMatrix(m2)
	PrintMatrix(Transpose(m2))
}
