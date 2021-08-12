package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process("https://whatcms.org/API/Tech?key=8291b3042eaf6d82b37f6af754422e479f2e454a7f094e0b8d8e978636806b79f54943&url=https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/")
	}
}
