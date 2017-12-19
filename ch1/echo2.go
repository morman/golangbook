// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"testing"
)

func Benchmark2Main(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 1; i < len(os.Args); i++ {
			s, sep := "", ""
			for _, arg := range os.Args[1:] {
				s += sep + arg
				sep = " "
			}
			fmt.Println(s)
		}
	}
}
