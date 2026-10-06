package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	size := 3
	arr := generateRandomElements(size)
	assert.Equal(t, size, len(arr))
}

func TestZeroSizeOnGenerateRandomElements(t *testing.T) {
	size := 0
	arr := generateRandomElements(size)
	assert.Equal(t, size, len(arr))
}

func TestMaximum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, 7, maximum(arr))
}

func TestMaximumEmptyArr(t *testing.T) {
	var arr []int
	assert.Equal(t, 0, maximum(arr))
}

func TestMaximumOneElement(t *testing.T) {
	arr := []int{1}
	assert.Equal(t, 1, maximum(arr))
}
