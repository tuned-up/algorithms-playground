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
	var cellCount int

	flag.IntVar(&cellCount, "cells count", 4, "Cells count in linked list")
	flag.Parse()

	if cellCount < 1 {
		fmt.Println("Please, create linked list with cell count > 0")
		return
	}

	ll := createLinkedList(cellCount)
	fmt.Println("Generated linked list: ")
	ll.Print()

	newList := reverse(ll)
	newList.Print()
	ll = reverse(newList)
	ll.Print()

	sorted := sortLinkedList(ll)
	fmt.Println("Sorted linked list: ")
	sorted.Print()

	//probably free ll
}

type linkedListNode struct {
	value int
	next *linkedListNode
}

func (ll *linkedListNode) Print() {
	runner := ll
	for runner != nil {
		if runner.next != nil {
			fmt.Printf("%d->", runner.value)
		} else {
			fmt.Printf("%d", runner.value)
		}
		runner = runner.next
	}
	fmt.Println()
}

func createLinkedList(cellCount int) *linkedListNode {
	rand.Seed(time.Now().Unix())
	ll := &linkedListNode{
		value: randValue(),
	}
	bottom := ll
	for i := 1; i < cellCount; i++ {
		newNode := &linkedListNode{value: randValue()}
		bottom.next = newNode
		bottom = newNode
	}

	return ll
}

func sortLinkedList(ll *linkedListNode) *linkedListNode {
	sorted := &linkedListNode{value: ll.value}
	ll = ll.next
	for ll != nil {
		runner := sorted
		prev := sorted
		for runner != nil && runner.value < ll.value  {
			prev = runner
			runner = runner.next
		}
		newNode := &linkedListNode{value: ll.value}
		if prev == runner {
			newNode.next = sorted
			sorted = newNode
		} else {
			prev.next = newNode
			newNode.next = runner
		}

		ll = ll.next
	}

	return sorted
}

func reverse(ll *linkedListNode) *linkedListNode {
	var prev *linkedListNode
	curr := ll

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func randValue() int {
	return rand.Intn(maxRand - minRand) + minRand
}