package main

import (
	"bytes"
	"fmt"
	"go/format"
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
		// The previous subsets are not actually stored again.
		// This has the disadvantage of not being able to hold the subsets in order.
		subsets[i] = make([][]uint8, 0)
		for j := 0; j < i; j++ {
			for k := range subsets[j] {
				buffer := make([]byte, len(prev)+1)
				copy(buffer, prev)
				buffer[len(buffer)-1] = uint8(i)
				subsets[i] = append(subsets[i], buffer)
			}
		}
	}

	buff := new(bytes.Buffer)
	fmt.Fprintf(buff, "package subsets\n\n")
	fmt.Fprintf(buff, "// Automatically generated - do not edit\n")
	fmt.Fprintf(buff, "// Limit = %d\n", limit)

	fmt.Fprintf(buff, "\nvar subsets = [][][]uint8{\n")
	for i := range subsets {
		fmt.Fprintf(buff, "\t{ ")
		for j := range subsets[i] {
			fmt.Fprintf(buff, "{")
			for k := range subsets[i][j] {
				fmt.Fprint(buff, subsets[i][j][k], ", ")
			}
			fmt.Fprintf(buff, "}, \n")
		}
		fmt.Fprintln(buff, "},")
	}

	fmt.Fprintf(buff, "}\n")

	formatted, err := format.Source(buff.Bytes())
	if err != nil {
		// Must be a coding error
		panic(err)
	}

	f, err := os.Create("../subsets-generated.go")
	if err == nil {
		_, err = f.Write(formatted)
	}
	if err == nil {
		err = f.Close()
	}

	if err != nil {
		panic(err)
	}
}
