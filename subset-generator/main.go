package main

import (
	"flag"
	"fmt"
)

var (
	// Max number of items to generate lookup table for. Since the output
	// is O(2^N), this can't be too high!
	limit int
)

// Run subset-generator to produce the boiler-plate to be copied to the end of subsets.go
func main() {
	flag.IntVar(&limit, "limit", 16, "Maximum value of 'n' to be cached (can not exceed 16)")
	flag.Parse()

	fmt.Println("const (")
	fmt.Printf("	limit = %d\n", limit)
	fmt.Print("	subsets = `")

	max := uint64(1<<uint(limit)) - 1
	for bits := uint64(1); bits <= max; bits++ {
		for bit := 0; bit <= limit; bit++ {
			if (bits & (1 << uint(bit))) != 0 {
				fmt.Printf("%c", "ABCDEFGHIJKLMNOPQRSTUVXYZ"[bit])
			}
		}
		fmt.Println()
	}

	fmt.Printf("`\n)\n")
}
