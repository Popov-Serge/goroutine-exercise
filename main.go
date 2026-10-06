package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size)
	}

	return arr
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxInt := data[0]
	for _, v := range data[1:] {
		if v > maxInt {
			maxInt = v
		}
	}

	return maxInt
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup

	if len(data) == 0 {
		return 0
	}
	chunks := min(CHUNKS, len(data))
	maxArr := make([]int, chunks)
	chunkSize := len(data) / chunks
	wg.Add(chunks)
	for i := range chunks {
		slice := data[chunkSize*i : chunkSize*(i+1)]
		go func(index int, arr []int) {
			defer wg.Done()
			maxArr[index] = maximum(arr)
		}(i, slice)
	}
	wg.Wait()
	return maximum(maxArr)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	arr := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxOne := maximum(arr)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxOne, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxTwo := maxChunks(arr)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxTwo, elapsed)
}
