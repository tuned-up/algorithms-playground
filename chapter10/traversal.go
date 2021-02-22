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

type BinaryTree struct {
	topNode *TreeNode
	maxLevel, elemCount int
}

type TreeNode struct {
	value int
	lChild, rChild *TreeNode
}

func (tn *TreeNode) hasChildren() bool {
	return tn.lChild != nil || tn.rChild != nil
}

func (bt *BinaryTree) AddNode(newValue int) {
	bt.elemCount++

	if bt.topNode == nil {
		bt.topNode = &TreeNode{value: newValue}
		bt.maxLevel++
		return
	}

	runner := bt.topNode
	level := 1
	for {
		level++
		if level > bt.maxLevel {
			bt.maxLevel = level
		}
		if runner.value > newValue {
			if runner.lChild == nil {
				runner.lChild = &TreeNode{value: newValue}
				return
			}
			runner = runner.lChild
		} else {
			if runner.rChild == nil {
				runner.rChild = &TreeNode{value: newValue}
				return
			}
			runner = runner.rChild
		}
	}
}

func (bt *BinaryTree) PrintTree() {
	bt.topNode.print("", "")
}

func (tn *TreeNode) print(prefix, childPrefix string) {
	fmt.Print(prefix)
	fmt.Println(tn.value)
	sep := "└── "
	link := "    "

	if tn.rChild != nil {
		sibling := tn.lChild != nil
		if sibling {
			sep = "├──"
			link = "|   "
		}
		next := tn.rChild
		if next.hasChildren() {
			next.print(childPrefix + sep, childPrefix + link)
		} else {
			next.print(childPrefix + sep, childPrefix + "    ")
		}
	}
	if tn.lChild != nil {
		next := tn.lChild
		if next.hasChildren() {
			next.print(childPrefix + "└── ", childPrefix + "    ")
		} else {
			next.print(childPrefix + "└── ", childPrefix + "    ")
		}
	}
}

func (bt *BinaryTree) TraverseSym() {
	bt.topNode.traverseNodeSym()
	fmt.Println()
}

func (tn *TreeNode) traverseNodeSym() {
	if tn.lChild != nil {
		tn.lChild.traverseNodeSym()
	}
	fmt.Printf("%d | ", tn.value)
	if tn.rChild != nil {
		tn.rChild.traverseNodeSym()
	}
}

func (bt *BinaryTree) TraverseReverse() {
	bt.topNode.traverseNodeReverse()
	fmt.Println()
}

func (tn *TreeNode) traverseNodeReverse() {
	if tn.lChild != nil {
		tn.lChild.traverseNodeSym()
	}
	if tn.rChild != nil {
		tn.rChild.traverseNodeSym()
	}
	fmt.Printf("%d | ", tn.value)
}

func (bt *BinaryTree) TraverseWidth() {
	bt.topNode.traverseNodeWidth(bt.elemCount)
	fmt.Println()
}

func (tn *TreeNode) traverseNodeWidth(elemCount int) {
	queue := make([]*TreeNode, 0, elemCount)
	queue = append(queue, tn)
	for i := 0; i < len(queue); i ++ {
		if queue[i] == nil {
			return
		}
		node := queue[i]
		fmt.Printf("%d | ", node.value)
		if node.lChild != nil {
			queue = append(queue, node.lChild)
		}
		if node.rChild != nil {
			queue = append(queue, node.rChild)
		}
	}
}

func main()  {
	var elemCount int

	flag.IntVar(&elemCount, "elements count", 15, "Elements count in tree")
	flag.Parse()

	randomArray := createRandomArray(elemCount)

	fmt.Printf("Unsorted: %v\n", randomArray)
	bt := &BinaryTree{}
	for _, v := range randomArray {
		bt.AddNode(v)
	}
	bt.PrintTree()

	bt.TraverseSym()
	bt.TraverseReverse()
	bt.TraverseWidth()
}

func createRandomArray(count int) []int {
	rand.Seed(time.Now().Unix())
	arr := make([]int, count, count)
	for i := 0; i < count; i++ {
		arr[i] = rand.Intn(maxRand - minRand) + minRand
	}

	return arr
}
