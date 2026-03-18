package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	arr := make([]int, 1000000000)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}

	if len(os.Args) == 2 {
		amountThreads, _ := strconv.Atoi(os.Args[1])
		start1 := time.Now()
		AsyncSum(arr, amountThreads)
		fmt.Println("-------->", time.Since(start1))
	} else {
		os.Exit(1)
	}
	start2 := time.Now()
	SyncSum(arr)
	fmt.Println("-------->", time.Since(start2))
}

func SyncSum(arr []int) {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	fmt.Println("Sync SUM: ")
	fmt.Println(sum)
}

func AsyncSum(arr []int, amountThreads int) {
}
