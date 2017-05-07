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
	subsets = make([][][]uint8, limit+1)
)

func main() {
	// subsets([]) == []
	subsets[0] = [][]uint8{buffer[0:0]}
	// subsets([1]) == [[1]]
	buffer = append(buffer, 1)
	subsets[1] = [][]uint8{[]uint8{}, buffer[0:1]}
	fmt.Println(subsets)

	for i := 2; i <= limit; i++ {
		// subsets([0..N]) = subsets([0..N-1]) ++ eachN(subsets([0..N-1]) ++ [N])
		subsets[i] = make([][]uint8, 0)
		for j := range subsets[i-1] {
			subsets[i] = append(subsets[i], subsets[j]...)
		}
		for _, prev := range subsets[i-1] {
			l := len(buffer)
			buffer = append(buffer, prev...)
			buffer = append(buffer, uint8(i))
			subsets[i] = append(subsets[i], buffer[l:])
		}
		fmt.Println("**", i, subsets[i])
	}
	fmt.Println(buffer)
}
