package main

import (
	"fmt"
	"os"
)

const (
	// Max number of items to generate lookup table for. Since the output
	// is O(2^N), this can't be too high!
	limit = 14
)

var (
	subsets = make([][][]uint8, limit+1)
)

func main() {
	// subsets([]) == []
	subsets[0] = [][]uint8{nil}

	for i := 1; i <= limit; i++ {
		// subsets([0..N]) = subsets([0..N-1]) ++ eachN(subsets([0..N-1]) ++ [N])
		// The previous subsets are not actually stored again.
		subsets[i] = make([][]uint8, 0)
		for j := 0; j < i; j++ {
			for _, prev := range subsets[j] {
				buffer := make([]byte, len(prev)+1)
				copy(buffer, prev)
				buffer[len(buffer)-1] = uint8(i)
				subsets[i] = append(subsets[i], buffer)
			}
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
		// /*  2 */ [][]uint8{[]uint8{0x2}, []uint8{0x1, 0x2}},
		// /*  2 */ [][]uint8{{2}, {1, 2}},
		fmt.Fprintf(f, "\t{ ")
		for j := range subsets[i] {
			fmt.Fprintf(f, "{")
			for k := range subsets[i][j] {
				fmt.Fprint(f, subsets[i][j][k], ", ")
			}
			fmt.Fprintf(f, "}, \n")
		}
		fmt.Fprintln(f, "},")
	}

	fmt.Fprintf(f, "}\n")
}
