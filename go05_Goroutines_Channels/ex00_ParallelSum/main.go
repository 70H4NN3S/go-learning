package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
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

func AsyncSum(arr []int, threads int) {
	size := len(arr) / threads
	ch := make(chan int, threads)
	var wg sync.WaitGroup
	sum := 0
	i := 0
	for ; i < threads-1; i++ {
		wg.Add(1)
		go func(subArr []int) {
			defer wg.Done()
			s := 0
			for _, n := range subArr {
				s += n
			}
			ch <- s
		}(arr[i*size : (i+1)*size])
	}
	wg.Add(1)
	go func(subArr []int) {
		defer wg.Done()
		s := 0
		for _, n := range subArr {
			s += n
		}
		ch <- s
	}(arr[i*size:])

	go func() {
		wg.Wait()
		close(ch)
	}()
	for pSum := range ch {
		sum += pSum
	}
	fmt.Println("Async SUM: ")
	fmt.Println(sum)
}
