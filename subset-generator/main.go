package main

import (
	"fmt"
)

const (
	// Max number of items to generate lookup table for. Since the output
	// is O(2^N), this can't be too high!
	//limit = 16
	limit = 5
)

var (
	buffer  = make([]uint8, 0, 1<<(limit+1))
	subsets = make([][]struct{ start, end int }, limit+1)
)

func main() {
	// subsets([]) == []
	subsets[0] = []struct{ start, end int }{{0, 0}}

	for i := 1; i <= limit; i++ {
		// subsets([0..N]) = subsets([0..N-1]) ++ eachN(subsets([0..N-1]) ++ [N])
		subsets[i] = make([]struct{ start, end int }, 0)
		for j := range subsets[i-1] {
			subsets[i] = append(subsets[i], subsets[i-1][j])
		}
		for _, prev := range subsets[i-1] {
			l := len(buffer)
			buffer = append(buffer, buffer[prev.start:prev.end]...)
			buffer = append(buffer, uint8(i))
			subsets[i] = append(subsets[i], struct{ start, end int }{l, len(buffer)})
		}
	}

	// TODO -could compress buffer much more: var buffer = []byte{0x1, 0x2, 0x1, 0x2, 0x3, ...
	fmt.Printf(`package subsets
// Buffer is %d bytes.

var buffer = %#v
var subsets = [][][]uint8{
`, len(buffer), buffer)

	for i := range subsets {
		fmt.Print("\t[][]uint8{")
		for j := range subsets[i] {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("buffer[%d:%d]", subsets[i][j].start, subsets[i][j].end)
		}
		fmt.Println("},")
	}

	fmt.Println("}")
}
