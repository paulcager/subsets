package main

import (
	"fmt"
	"os"
)

const (
	// Max number of items to generate lookup table for. Since the output
	// is O(2^N), this can't be too high!
	limit = 16
)

var (
	subsets = make([][][]uint8, limit+1)
)

func main() {
	// subsets([]) == []
	subsets[0] = [][]uint8{nil}

	for i := 1; i <= limit; i++ {
		// subsets([0..N]) = subsets([0..N-1]) ++ eachN(subsets([0..N-1]) ++ [N])
		subsets[i] = make([][]uint8, 0)
		for j := range subsets[i-1] {
			subsets[i] = append(subsets[i], subsets[i-1][j])
		}
		for _, prev := range subsets[i-1] {
			buffer := make([]byte, len(prev)+1)
			copy(buffer, prev)
			buffer[len(buffer)-1] = uint8(i)
			subsets[i] = append(subsets[i], buffer)
		}
	}

	f, err := os.Create("../subsets-generated.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "package subsets\n\n")
	fmt.Fprintf(f, "// Automatically generated - do not edit\n")
	fmt.Fprintf(f, "// Limit = %d\n", limit)

	fmt.Fprintf(f, "\nvar subsets = [][][]uint8{\n")
	for i := range subsets {
		fmt.Fprintf(f, "\t/* % 2d */\t%#v,\n", i, subsets[i])
	}

	fmt.Fprintf(f, "}\n")
}
