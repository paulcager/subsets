package subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThree(t *testing.T) {
	returned := make(map[string]bool)
	callback := func(indexes []int) {
		b := make([]byte, len(indexes))
		for i := range indexes {
			b[i] = byte('A' + indexes[i])
		}
		returned[string(b)] = true
	}
	Enumerate(3, callback)
	expected := map[string]bool{
		"A":   true,
		"B":   true,
		"C":   true,
		"AB":  true,
		"AC":  true,
		"BC":  true,
		"ABC": true,
	}
	assert.Equal(t, expected, returned)
}

func nullCallback(_ []int) {}
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
