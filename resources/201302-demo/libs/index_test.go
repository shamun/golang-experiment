package main

import (
    "testing"
    "strings"
)

func TestIndex(t *testing.T) {
    var tests = []struct {
        s string
        sep string
        out int
    }{
        {"", "", 0},
        {"fo", "foo", -1},
        {"foo", "foo", 0},
    }
    for _, test := range tests {
        actual := strings.Index(test.s, test.sep)
        if actual != test.out {
            t.Errorf("Index(%q.%q) = $v; want %v", test.s, test.sep, actual, test.out)
        }
    }
}
