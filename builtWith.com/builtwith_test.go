package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process("https://api.builtwith.com/v19/api.json?KEY=4d94c975-7c4c-4a57-b8d9-9091acfdc4ac&LOOKUP=https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/")
	}
}
