package main

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {

	var sum int64
	var wg sync.WaitGroup

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	max := len(numbers)

	sizeOfParts := max / numCPU

	for i := 0; i < numCPU; i++ {
		start := i * sizeOfParts
		end := start + sizeOfParts
		part := numbers[start:end]
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			var partSum int64

			for _, n := range nums {
				partSum += int64(n)
			}
			sum += partSum
		}(part)
		wg.Wait()
	}

	return sum
}
