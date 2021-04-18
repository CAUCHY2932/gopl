package main

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse()
	}
}
