package main

import (
    "strings"
    "testing"
)

func BenchmarkIndex(b *testing.B) {
    const s = "some_text=some_value"
    for i := 0; i < b.N; i++ {
        strings.Index(s, "v")
    }
}
