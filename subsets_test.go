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
	t[string(b)] = true
}

func TestThree(t *testing.T) {
	callback := make(testCallback)
	Enumerate(3, callback.callback)
	expected := map[string]bool{
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
	callback1 := make(testCallback)
	Enumerate(8, callback1.callback)
	callback2 := make(testCallback)
	traditionalEnumerate(8, callback2.callback)
	assert.Equal(t, callback2, callback1)
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

func traditionalEnumerate(n int, callback Callback) {
	max := uint64(1<<uint(n)) - 1
	indexes := make([]int, n)
	for bits := uint64(1); bits <= max; bits++ {
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
