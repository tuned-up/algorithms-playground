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

	flag.IntVar(&elemCount, "elements count", 5, "Elements count in array list")
	flag.Parse()

	randomArray := createRandomArray(elemCount)

	fmt.Printf("Unsorted: %v\n", randomArray)
	qcksrt(randomArray)
	fmt.Printf("Sorted: %v\n", randomArray)
}

// so quick it skips vowels
func qcksrt(array []int) {
	start, end := 0, len(array) - 1
	if start >= end {
		return
	}

	divider := array[start]

	low, high := start, end
	for {
		for array[high] >= divider && high > low {
			high--
		}
		if high <= low {
			array[low] = divider
			break
		}

		tmp := array[low]
		array[low] = array[high]
		array[high] = tmp

		low++
		for array[low] <= divider && low < high {
			low++
		}
		if low >= high {
			low = high
			array[low] = divider
			break
		}

		tmp = array[high]
		array[high] = array[low]
		array[low] = tmp
	}

	qcksrt(array[start:(low+1)])
	qcksrt(array[(low+1):(end+1)])
}

func qcksrt_elegant(array []int) {
	start, end := 0, len(array) - 1
	if start >= end {
		return
	}

	divider := array[start]

	low, high := start, end
	for {
		for array[low] < divider && low <= end {
			low++
		}
		for array[high] > divider && high >= start {
			high--
		}

		if low >= high {
			break
		}

		tmp := array[high]
		array[high] = array[low]
		array[low] = tmp
	}

	qcksrt(array[start:(low+1)])
	qcksrt(array[(low+1):(end+1)])
}

func createRandomArray(count int) []int {
	rand.Seed(time.Now().Unix())
	arr := make([]int, count, count)
	for i := 0; i < count; i++ {
		arr[i] = rand.Intn(maxRand - minRand) + minRand
	}

	return arr
}


