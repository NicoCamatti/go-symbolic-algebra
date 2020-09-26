package symbolic

// A Node in a binary tree
type node struct {
	left  Evaluatable
	right Evaluatable
}

func (n *node) FunctionOf(v *Variable) bool {
	leftIsNil := n.left == nil
	rightIsNil := n.right == nil
	if !leftIsNil && !rightIsNil {
		return n.left.FunctionOf(v) || n.right.FunctionOf(v)
	} else if !leftIsNil {
		return n.left.FunctionOf(v)
	} else if !rightIsNil {
		return n.right.FunctionOf(v)
	} else {
		return false
	}
}

func (n *node) IsConstant() bool {
	leftIsNil := n.left == nil
	rightIsNil := n.right == nil
	if !leftIsNil && !rightIsNil {
		return n.left.IsConstant() && n.right.IsConstant()
	} else if !leftIsNil {
		return n.left.IsConstant()
	} else if !rightIsNil {
		return n.right.IsConstant()
	} else {
		return true
	}
}
