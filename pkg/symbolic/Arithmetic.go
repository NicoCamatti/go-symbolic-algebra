package symbolic

// add operation node
type add struct {
	node
}

// Returns a Add Node given the left and right operands: left + rights
func NodeAdd(left, right Evaluatable) Evaluatable {
	parent := node{left: left, right: right}
	return &add{parent}
}

func (a *add) Evaluate() float64 {
	return a.left.Evaluate() + a.right.Evaluate()
}

func (a *add) Diff(v *Variable) Evaluatable {
	leftIsFunc := a.left.FunctionOf(v)
	rightIsFunc := a.right.FunctionOf(v)
	if leftIsFunc && rightIsFunc {
		return NodeAdd(a.left.Diff(v), a.right.Diff(v))
	} else if leftIsFunc {
		return a.left.Diff(v)
	} else if rightIsFunc {
		return a.right.Diff(v)
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (a *add) String() string {
	return "(" + a.left.String() + " + " + a.right.String() + ")"
}

// sub operation node
type sub struct {
	node
}

// Returns a Sub Node given the left and right operands: left - right
func NodeSub(left, right Evaluatable) Evaluatable {
	parent := node{left: left, right: right}
	return &sub{parent}
}

func (s *sub) Evaluate() float64 {
	return s.left.Evaluate() - s.right.Evaluate()
}

func (s *sub) Diff(v *Variable) Evaluatable {
	leftIsFunc := s.left.FunctionOf(v)
	rightIsFunc := s.right.FunctionOf(v)
	if leftIsFunc && rightIsFunc {
		return NodeSub(s.left.Diff(v), s.right.Diff(v))
	} else if leftIsFunc {
		return s.left.Diff(v)
	} else if rightIsFunc {
		return NodeMultiply(GetConstant(CONSTANT_MINUS_ONE), s.right.Diff(v))
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (s *sub) String() string {
	return "(" + s.left.String() + " - " + s.right.String() + ")"
}

// multiply operation node
type multiply struct {
	node
}

// Returns a Multiply Node given the left and right operands: left * right
func NodeMultiply(left, right Evaluatable) Evaluatable {
	parent := node{left: left, right: right}
	return &multiply{parent}
}

func (m *multiply) Evaluate() float64 {
	return m.left.Evaluate() * m.right.Evaluate()
}

func (m *multiply) Diff(v *Variable) Evaluatable {
	leftIsFunc := m.left.FunctionOf(v)
	rightIsFunc := m.right.FunctionOf(v)
	if leftIsFunc && rightIsFunc {
		return NodeAdd(
			NodeMultiply(m.left.Diff(v), m.right),
			NodeMultiply(m.left, m.right.Diff(v)),
		)
	} else if leftIsFunc {
		return NodeMultiply(m.left.Diff(v), m.right)
	} else if rightIsFunc {
		return NodeMultiply(m.left, m.right.Diff(v))
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (m *multiply) String() string {
	return "(" + m.left.String() + " * " + m.right.String() + ")"
}

// divide operation Node
type divide struct {
	node
}

// Returns a Divide Node given the left and right operands: left / right
func NodeDivide(left, right Evaluatable) Evaluatable {
	parent := node{left: left, right: right}
	return &divide{parent}
}

func (d *divide) Evaluate() float64 {
	denominator := d.right.Evaluate()
	if denominator == GetConstant(CONSTANT_ZERO).Evaluate() {
		panic("Division by zero occured!")
	} else {
		return d.left.Evaluate() / denominator
	}
}

func (d *divide) Diff(v *Variable) Evaluatable {
	leftIsFunc := d.left.FunctionOf(v)
	rightIsFunc := d.right.FunctionOf(v)
	if leftIsFunc && rightIsFunc {
		return NodeDivide(
			NodeSub(
				NodeMultiply(d.right, d.left.Diff(v)),
				NodeMultiply(d.left, d.right.Diff(v)),
			),
			NodeMultiply(d.right, d.right),
		)
	} else if leftIsFunc {
		return NodeDivide(d.left.Diff(v), d.right)
	} else if rightIsFunc {
		return NodeDivide(
			NodeMultiply(
				NodeMultiply(GetConstant(CONSTANT_MINUS_ONE), d.left),
				d.right.Diff(v),
			),
			NodeMultiply(d.right, d.right),
		)
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (d *divide) String() string {
	return "(" + d.left.String() + " / " + d.right.String() + ")"
}
