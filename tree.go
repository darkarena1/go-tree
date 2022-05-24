// Package tree provides an implementation of a binary tree with threaded channels that help
// traverse the tree much faster, taking advantage of go's threaded nature.
package tree

// Node provides the implementation of the tree node
type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// SubtreeChannel provides a channel where each node in the tree
// starting with this node will be returned
func (n Node[T]) SubtreeChannel() chan *Node[T] {
	channel := make(chan *Node[T], 50)
	go n.addNodeToChannel(channel)
	return channel
}

// TerminalElementsChannel provides a channel that contains all
// terminal nodes (that is nodes that don't have either a left
// or right node
func (n Node[T]) TerminalElementsChannel() chan *Node[T] {
	channel := make(chan *Node[T], 50)
	go n.findTerminalElements(channel)
	return channel
}

// findTerminalElements adds all the terminal elements from the
// current node to the channel, and closes the channel
func (n Node[T]) findTerminalElements(channel chan *Node[T]) {
	defer close(channel)
	n.addTerminalElementsToChannel(channel)
}

// addTerminalElementsToChannel adds all terminate elements from the
// current node to the channel but DOES NOT close the channel.  This
// method is recursive.
func (n Node[T]) addTerminalElementsToChannel(channel chan *Node[T]) {
	nodePtr := &n
	if nodePtr.Left != nil {
		nodePtr.Left.addTerminalElementsToChannel(channel)
	}

	if nodePtr.Right != nil {
		nodePtr.Right.addTerminalElementsToChannel(channel)
	}

	//This is a terminal element
	if nodePtr.Left == nil && nodePtr.Right == nil {
		channel <- nodePtr
	}
}

// addNodeToChannel adds each node it encounters to the channel and
// closes the channel.  The order of these additions will be the same
// as the distance of the root node going from left to right.
func (n Node[T]) addNodeToChannel(channel chan *Node[T]) {
	defer close(channel)
	nodePtr := &n
	toBeProcessed := []*Node[T]{nodePtr}
	channel <- nodePtr
	for len(toBeProcessed) > 0 {
		nodePtr, toBeProcessed = toBeProcessed[0], toBeProcessed[1:]
		if nodePtr.Left != nil {
			channel <- nodePtr.Left
			toBeProcessed = append(toBeProcessed, nodePtr.Left)
		}

		if nodePtr.Right != nil {
			channel <- nodePtr.Right
			toBeProcessed = append(toBeProcessed, nodePtr.Right)
		}
	}

}
