package main

import (
	"flag"
	"fmt"
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

	randomArray := createRandomArray(elemCount)

	fmt.Printf("Unsorted: %v\n", randomArray)
	merge_sort(randomArray)
	fmt.Printf("Sorted: %v\n", randomArray)
}

func merge_sort(array []int) {
	arrayLength := len(array)
	mrg_srt(array, make([]int, arrayLength, arrayLength), 0, arrayLength - 1)
}

func mrg_srt(array, temp []int, start, end int) {
	if start == end {
		return
	}

	midpoint := (start + end) / 2
	mrg_srt(array, temp, start, midpoint)
	mrg_srt(array, temp, midpoint + 1, end)

	leftIndex := start
	rightIndex := midpoint + 1
	runner := leftIndex
	for (leftIndex <= midpoint) && (rightIndex <= end) {
		if array[leftIndex] <= array[rightIndex] {
			temp[runner] = array[leftIndex]
			leftIndex++
		} else {
			temp[runner] = array[rightIndex]
			rightIndex++
		}
		runner++
	}

	for i := leftIndex; i <= midpoint; i++ {
		temp[runner] = array[i]
		runner++
	}
	for i := rightIndex; i <= end; i++ {
		temp[runner] = array[i]
		runner++
	}

	for i := 0; i <= end; i++ {
		array[i] = temp[i]
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
