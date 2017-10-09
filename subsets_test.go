package subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func nullCallback(_ []int) {}

type testCallback map[string]bool

func (t testCallback) callback(indexes []int) {
	b := make([]byte, len(indexes))
	for i := range indexes {
		b[i] = byte('A' + indexes[i])
	}
	if t[string(b)] {
		panic("Duplicate: " + string(b))
	}
	t[string(b)] = true
}

func traditionalEnumerate(n int, callback Callback) {
	max := uint64(1<<uint(n)) - 1
	indexes := make([]int, n)
	for bits := uint64(0); bits <= max; bits++ {
		size := 0
		for bit := 0; bit <= n; bit++ {
			if (bits & (1 << uint(bit))) != 0 {
				indexes[size] = bit
				size++
			}
		}
		callback(indexes[:size])
	}
}

func recursiveManiacsMethod(n int) [][]int {
	if n == 0 {
		return [][]int{[]int{}}
	}
	result := recursiveManiacsMethod(n - 1)
	prev := len(result)
	for i := 0; i < prev; i++ {
		newOne := make([]int, len(result[i])+1)
		copy(newOne, result[i])
		newOne[len(newOne)-1] = n - 1
		result = append(result, newOne)
	}

	return result
}

func TestThree(t *testing.T) {
	callback := make(testCallback)
	Enumerate(3, callback.callback)
	expected := map[string]bool{
		"":    true,
		"A":   true,
		"B":   true,
		"C":   true,
		"AB":  true,
		"AC":  true,
		"BC":  true,
		"ABC": true,
	}
	assert.Equal(t, expected, map[string]bool(callback))
}

func TestCompare(t *testing.T) {
	const n = 3

	callback1 := make(testCallback)
	traditionalEnumerate(n, callback1.callback)

	callback2 := make(testCallback)
	Enumerate(n, callback2.callback)

	callback3 := make(testCallback)
	items := recursiveManiacsMethod(n)
	for _, item := range items {
		b := make([]byte, len(item))
		for i := range item {
			b[i] = byte('A' + item[i])
		}
		callback3[string(b)] = true
	}

	assert.Equal(t, callback1, callback2)
	assert.Equal(t, callback1, callback3)
}

func BenchmarkThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Enumerate(3, nullCallback)
	}
}
func BenchmarkSeven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Enumerate(7, nullCallback)
	}
}
func BenchmarkFourteen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Enumerate(14, nullCallback)
	}
}
func BenchmarkSixteen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Enumerate(16, nullCallback)
	}
}

func BenchmarkThreeTrad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		traditionalEnumerate(3, nullCallback)
	}
}
func BenchmarkSixteenTrad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		traditionalEnumerate(16, nullCallback)
	}
}

var sink int

func BenchmarkSixteenRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += len(recursiveManiacsMethod(16))
	}
}

func BenchmarkSixteenIsItReallyWorthIt(b *testing.B) {
	// Do a little bit of work for each iteration so I can get a feel for how significant
	// the cost saving is likely to be.
	var sink int
	callbackWithWork := func(indexes []int) {
		for i := range indexes {
			sink = sink<<7 + indexes[i]
			sink = sink<<4 + i
		}
	}
	for i := 0; i < b.N; i++ {
		Enumerate(16, callbackWithWork)
	}
}
