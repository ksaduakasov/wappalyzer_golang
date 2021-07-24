package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := GenerateNumbers(1e7)

	t := time.Now()
	sum := Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))

	t = time.Now()
	sum = AddConcurrent(numbers)
	fmt.Printf("Concurrent Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
}