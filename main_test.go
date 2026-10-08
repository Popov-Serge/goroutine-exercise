package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "positive size",
			size: 3,
			want: 3,
		},
		{
			name: "zero size",
			size: 0,
			want: 0,
		},
		{
			name: "negative size",
			size: -1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := generateRandomElements(tt.size)
			assert.Equal(t, tt.want, len(arr))
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "normal",
			data: []int{1, 2, 3, 4, 5, 6, 7},
			want: 7,
		},
		{
			name: "empty",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maximum(tt.data))
		})
	}
}
