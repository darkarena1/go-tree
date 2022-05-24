package tree

import (
	"testing"
)

func TestNode_Left(t *testing.T) {
	root := Node[int]{Value: 1}
	leftLeaf := Node[int]{Value: 2}
	root.Left = &leftLeaf

	if root.Left.Value != 2 {
		t.Failed()
	}
}

func TestNode_SubtreeChannel(t *testing.T) {
	nodes := make([]Node[int], 100)
	for i := 0; i < 100; i++ {
		nodes[i] = Node[int]{Value: i}
	}

	//Use fiz buzz algorithm to fill out tree,  Terminal nodes are intersection of 3 & 5
	currentParent := 1
	currentChild := 3
	nodes[0].Left = &nodes[1]
	nodes[0].Right = &nodes[2]

	for currentChild < 100 {

		if currentParent%3 != 0 {
			nodes[currentParent].Left = &nodes[currentChild]
			currentChild++
		}

		if currentParent%5 != 0 && currentChild < 100 {
			nodes[currentParent].Right = &nodes[currentChild]
			currentChild++
		}
		currentParent++
	}

	currentValue := 0
	for s := range nodes[0].SubtreeChannel() {
		if s.Value != currentValue {
			t.Failed()
		}
		currentValue++
	}

}

func TestNode_TerminalElementsChannel(t *testing.T) {
	nodes := make([]Node[int], 100)
	for i := 0; i < 100; i++ {
		nodes[i] = Node[int]{Value: i}
	}

	//Use fiz buzz algorithm to fill out tree,  Terminal nodes are intersection of 3 & 5
	currentParent := 1
	currentChild := 3
	nodes[0].Left = &nodes[1]
	nodes[0].Right = &nodes[2]

	for currentChild < 100 {

		if currentParent%3 != 0 {
			nodes[currentParent].Left = &nodes[currentChild]
			currentChild++
		}

		if currentParent%5 != 0 && currentChild < 100 {
			nodes[currentParent].Right = &nodes[currentChild]
			currentChild++
		}
		currentParent++
	}

	currentChild = 67
	for s := range nodes[0].TerminalElementsChannel() {
		if !(s.Value > 66 || (s.Value%3 == 0 && s.Value%5 == 0)) {
			t.Failed()
		}
	}

}
