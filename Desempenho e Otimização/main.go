package main

import (
	"fmt"
	"time"
)

//func sumUpToN(N int) int {
//	sum := 0
//	for i := 1; i <= N; i++ {
//		sum += i
//	}
//	return sum
//}

func sumUpToN(N int) int {
	return (N * (N + 1)) / 2
}

func main() {
	N := 1000000

	start := time.Now()
	result := sumUpToN(N)
	fmt.Println("Sum:", result)
	fmt.Println("Time taken:", time.Since(start))
}
