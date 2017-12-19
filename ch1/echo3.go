// Compact version of echo
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Benchmark3Main(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
}
