package openapi

// PostorderTraversal traverses the tree in postorder.
func PostorderTraversal(root Node) []Node {
	if root == nil {
		return nil
	}

	var result []Node
	stack := &Stack{}
	stack.Push(root)

	for len(stack.data) > 0 {
		node := stack.Pop()

		result = append([]Node{node}, result...)

		// Push the children to the stack.
		for _, child := range node.Children() {
			stack.Push(child)
		}
	}

	return result
}
