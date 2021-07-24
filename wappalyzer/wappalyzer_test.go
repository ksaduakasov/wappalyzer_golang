package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process("https://api.wappalyzer.com/lookup/v2/?urls=https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang")
	}
}
