package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	minRand = 1
	maxRand = 100
)

func main()  {
	var elemCount int

	flag.IntVar(&elemCount, "elements count", 15, "Elements count in array list")
	flag.Parse()

	heap := createRandomArray(elemCount)

	fmt.Printf("Unsorted: %v\n", heap)
	heapsort(heap)
	fmt.Printf("Sorted: %v\n", heap)
}

func heapsort(heap []int) {
	makeBinaryTree(heap)
	fmt.Printf("Binary tree: %v\n", heap)

	for i := len(heap); i > 0; i-- {
		temp := heap[0]
		heap[0] = heap[i-1]
		heap[i-1] = temp
		repairTree(heap[:i-1])
	}
}

func makeBinaryTree(heap []int) {
	for i := 0; i < len(heap); i++ {
		index := i
		for index != 0 {
			parentIndex := (index - 1) / 2
			if heap[parentIndex] > heap[index] {
				break
			}

			tmp := heap[parentIndex]
			heap[parentIndex] = heap[index]
			heap[index] = tmp

			index = parentIndex
		}
	}
}

func repairTree(brokenTree []int) {
	maxIndex := len(brokenTree) - 1
	if maxIndex < 1 {
		return
	}

	index := 0
	for {
		lChild := 2 * index + 1
		rChild := 2 * index + 2

		if lChild > maxIndex {
			lChild = maxIndex
		}
		if rChild > maxIndex {
			rChild = maxIndex
		}

		if brokenTree[index] >= brokenTree[lChild] && brokenTree[index] >= brokenTree[rChild] {
			return
		}

		var swap int
		if brokenTree[lChild] > brokenTree[rChild] {
			swap = lChild
		} else {
			swap = rChild
		}

		temp := brokenTree[index]
		brokenTree[index] = brokenTree[swap]
		brokenTree[swap] = temp

		index = swap
	}
}

func easyPrintBinaryTree(heap []int) {
	heapLen := len(heap)
	if heapLen == 0 {
		return
	}

	fmt.Println(heap[0])
	i := 1
	runner := 0
	for {
		levelCount := int(math.Pow(2, float64(i)))

		for j := 0; j < levelCount; j++ {
			runner = runner + 1
			if runner >= heapLen {
				fmt.Println()

				return
			}
			fmt.Printf("%d | ", heap[runner])
		}
		fmt.Println()
		i++
	}
}

func createRandomArray(count int) []int {
	rand.Seed(time.Now().Unix())
	arr := make([]int, count, count)
	for i := 0; i < count; i++ {
		arr[i] = rand.Intn(maxRand - minRand) + minRand
	}

	return arr
}
