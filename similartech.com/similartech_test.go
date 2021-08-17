package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process("https://api.similartech.com/v1/site/wtotem.com/technologies?userkey=3b3b8cdb2b9b43dbad4a6138ceab60ea&format=json")
	}
}