package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process("https://api.similartech.com/v1/site/wtotem.com/technologies?userkey=59c73eb9ec2841e4b22396c36f27e6d2&format=json")
	}
}