package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	els := []int{6, 3, 2, 4, 5, 1}
	BubbleSort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, els[5])
	assert.EqualValues(t, 5, els[4])
	assert.EqualValues(t, 4, els[3])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 1, els[0])
}
func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}
func BenchmarkBubbleSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
func BenchmarkBubbleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
func BenchmarkBubbleSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
