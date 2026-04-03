// Package arrays is a simple package to test arrays and write a SUM function
package arrays

func Sum(arr []int) (res int) {
	for i := range arr {
		res += arr[i]
	}
	return
}

func SumAllInOne(numbersToSum ...[]int) (result int) {
	for _, slice := range numbersToSum {
		result += Sum(slice)
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	result := make([]int, len(numbersToSum))

	for i, slice := range numbersToSum {
		sum := Sum(slice)
		result[i] = sum
	}
	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	result := make([]int, len(numbersToSum))

	for i, slice := range numbersToSum {
		for j := 1; j < len(slice); j++ {
			result[i] += slice[j]
		}
	}
	return result
}
