package symbolic

// A Node in a binary tree
type node struct {
	left  Evaluatable
	right Evaluatable
}

func (n *node) FunctionOf(v *Variable) bool {
	return n.left.FunctionOf(v) || n.right.FunctionOf(v)
}
