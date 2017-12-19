// Compact version of echo
package main

import (
	"fmt"
	"os"
	"testing"
)

func Benchmark4Main(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for i := 1; i < len(os.Args); i++ {
			fmt.Printf("index: %d arg: %v \n", i, os.Args[i])
		}
	}
}
