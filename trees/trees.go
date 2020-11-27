package main

import (
	"fmt"
	"math"
)

// BST - binary search tree
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// TreeLevel - struct to store pointer to tree node and the level of it's parent
type TreeLevel struct {
	Node        *BST
	ParentLevel int
}

func main() {
	fmt.Println("init example tree")
	root := BST{Value: 10}
	root.insert(5)
	root.insert(2)
	root.insert(5)
	root.insert(1)
	root.insert(15)
	root.insert(13)
	root.insert(22)
	root.insert(14)
	// pretty print via level order
	root.levelOrderTraversal()
	//searching
	//search1 := root.search(22)
	//search2 := root.search(16)
	//search3 := root.search(1)
	//fmt.Printf("%v\t%v\t%v\n", search1, search2, search3)
	fmt.Printf("Max value: %d\n", root.maxValue())
	fmt.Printf("Min value: %d\n", root.minValue())
	//fmt.Printf("In Order traversal: %v\n", root.inOrder())
	//fmt.Printf("Post Order traversal: %v\n", root.postOrder())
	//fmt.Printf("Has path sum %d: %v\n", 20, root.hasPathSum(20))
	//fmt.Printf("Has path sum %d: %v\n", 18, root.hasPathSum(18))
	//fmt.Printf("Has path sum %d: %v\n", 52, root.hasPathSum(52))
	//fmt.Printf("Has path sum %d: %v\n", 51, root.hasPathSum(51))
	//fmt.Printf("Has path sum %d: %v\n", 6, root.hasPathSum(6))
	fmt.Printf("Tree is valid: %v\n", root.isValid())

}

// assumes a non-empty tree
func (tree *BST) insert(value int) {
	if tree.Value > value {
		if tree.Left != nil {
			tree.Left.insert(value)
		} else {
			tree.Left = &BST{Value: value}
		}
	} else {
		if tree.Right != nil {
			tree.Right.insert(value)
		} else {
			tree.Right = &BST{Value: value}
		}
	}
}

func (tree *BST) search(value int) bool {
	if tree.Value == value {
		return true
	}
	if tree.Value > value {
		if tree.Left != nil {
			return tree.Left.search(value)
		}
		return false
	}
	if tree.Right != nil {
		return tree.Right.search(value)
	}
	return false
}

func (tree *BST) maxValue() int {
	if tree.Right != nil {
		return tree.Right.maxValue()
	}
	return tree.Value
}

func (tree *BST) minValue() int {
	if tree.Left != nil {
		return tree.Left.minValue()
	}
	return tree.Value
}

// return in order traversal as an array
func (tree *BST) inOrder() []int {
	acc := []int{}
	if tree.Left != nil {
		acc = append(acc, tree.Left.inOrder()...)
	}
	acc = append(acc, tree.Value)
	if tree.Right != nil {
		acc = append(acc, tree.Right.inOrder()...)
	}
	return acc
}

// return post order (left right root) traversal as an array
func (tree *BST) postOrder() []int {
	acc := []int{}
	if tree.Left != nil {
		acc = append(acc, tree.Left.postOrder()...)
	}
	if tree.Right != nil {
		acc = append(acc, tree.Right.postOrder()...)
	}
	acc = append(acc, tree.Value)
	return acc
}

// if tree has a root-leaf path which has a given sum
func (tree *BST) hasPathSum(sum int) bool {
	if sum < tree.Value {
		return false
	}
	remainingSum := sum - tree.Value
	if (tree.Left == nil || tree.Right == nil) && remainingSum == 0 {
		return true
	}
	return (tree.Left != nil && tree.Left.hasPathSum(remainingSum)) ||
		(tree.Right != nil && tree.Right.hasPathSum(remainingSum))
}

// print all paths from root to leaf
func (tree *BST) printPaths() {

}

func (tree *BST) mirror() {

}

func (tree *BST) isValid() bool {
	return tree.validityHelper(math.MinInt32, math.MaxInt32)
}

func (tree *BST) validityHelper(minValue int, maxValue int) bool {
	if tree == nil {
		return true
	}
	return tree.Value >= minValue && tree.Value <= maxValue &&
		tree.Left.validityHelper(minValue, tree.Value) &&
		tree.Right.validityHelper(tree.Value, maxValue)
}

/**
Traverse tree in level order and print out node values
*/
func (tree *BST) levelOrderTraversal() {
	frontier, level := []TreeLevel{}, 0
	frontier = append(frontier, TreeLevel{Node: tree, ParentLevel: level})

	fmt.Println("Start level order traverse")
	previousLevel := 0
	for len(frontier) > 0 {
		treeLevel := frontier[0]
		frontier = frontier[1:]
		if treeLevel.Node.Left != nil {
			frontier = append(frontier, TreeLevel{Node: treeLevel.Node.Left, ParentLevel: treeLevel.ParentLevel + 1})
		}
		if treeLevel.Node.Right != nil {
			frontier = append(frontier, TreeLevel{Node: treeLevel.Node.Right, ParentLevel: treeLevel.ParentLevel + 1})
		}
		if treeLevel.ParentLevel != previousLevel {
			fmt.Printf("\n")
			previousLevel++
		}
		fmt.Printf("%d\t", treeLevel.Node.Value)
	}
	fmt.Printf("\n")
}

func (tree *BST) rightSideView() {
	// just implement level order traversal above and then for each level print the max value for that level
}

func (tree *BST) leftSideView() {
	// just implement level order traversal above and then for each level print the min value for that level
}
