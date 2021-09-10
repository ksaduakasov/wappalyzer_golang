package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process("https://api.larger.io/v1/search/key/G0KM2YWNWWAFMW6DZJ7Z66XPCE6ADOJE?domain=wtotem.com")
	}
}
