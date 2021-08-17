package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process("https://api.builtwith.com/v19/api.json?KEY=c4f0023c-b4d4-44b7-b5bd-4b0e1d4a8d2b&LOOKUP=https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/")
	}
}
