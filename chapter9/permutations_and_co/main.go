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

func main() {
	var elemCount int
	var permLength int

	flag.IntVar(&elemCount, "elements count", 5, "Elements count in array list")
	flag.IntVar(&permLength, "permutation length", 2, "Permutation length (k as in 'k from n')")
	flag.Parse()

	if permLength > elemCount {
		return
	}

	randomSet := createRandomArray(elemCount)
	fmt.Printf("%v\n", randomSet)

	perms := selectKfromNRecursive(permLength, randomSet)
	fmt.Printf("%v", perms)
}

func selectKfromNRecursive(k int, set []int) [][]int {
	res := make([][]int, 0, 0)
	doSelectKfromNRecursive(0, make([]int, k, k), set, &res)

	return res
}

func doSelectKfromNRecursive(level int, sel []int, set []int, res *[][]int) {
	if level == cap(sel) {
		result := make([]int, len(sel), cap(sel))
		for i := 0; i < len(sel); i++ {
			result[i] = set[sel[i]]
		}
		*res = append(*res, result)
	} else {
		var start int
		if level > 0 {
			start = sel[level - 1] + 1
		}
		for i := start; i < len(set); i++ {
			sel[level] = i

			doSelectKfromNRecursive(level+1, sel, set, res)
		}
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

