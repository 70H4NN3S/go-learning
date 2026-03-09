package main

func NewMatrix(rows, cols int) [][]float64 {
	m := make([][]float64, rows)
	for i := range m {
		m[i] = make([]float64, cols)

		// Fill with data to test
		for j := range m[i] {
			m[i][j] = float64((i*j + i + j) % 10)
		}
	}
	return m
}
