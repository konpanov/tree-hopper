package main

import (
	"fmt"
	"strings"
	"testing"
)

func assertIntEqual(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("%d != %d", a, b)
	}
}

func assertStringEqual(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("%s != %s", a, b)
	}
}

func assertByteEqual(t *testing.T, a, b byte) {
	if a != b {
		t.Errorf("%c != %c", a, b)
	}
}

func createNLines(n int) string {
	lines := make([]string, 0)
	for i := 0; i < n; i++ {
		lines = append(lines, fmt.Sprint(i))
	}
	return strings.Join(lines, "\r\n")
}
