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
		return GetConstant(ConstantZero)
	}
}

func (a *add) String() string {
	return "(" + a.left.String() + " + " + a.right.String() + ")"
}

func (a *add) Trim() Evaluatable {
	// Drop a sum with zero
	leftTrim := a.left.Trim()
	rightTrim := a.right.Trim()

	dropLeft := false
	if leftTrim.IsConstant() {
		if leftTrim.Evaluate() == 0.0 {
			dropLeft = true
		}
	}
	dropRight := false
	if rightTrim.IsConstant() {
		if rightTrim.Evaluate() == 0.0 {
			dropRight = true
		}
	}

	if dropLeft && dropRight {
		return GetConstant(ConstantZero)
	} else if dropLeft {
		return rightTrim
	} else if dropRight {
		return leftTrim
	} else {
		return NodeAdd(leftTrim, rightTrim)
	}
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
		return NodeMultiply(GetConstant(ConstantMinusOne), s.right.Diff(v))
	} else {
		return GetConstant(ConstantZero)
	}
}

func (s *sub) String() string {
	return "(" + s.left.String() + " - " + s.right.String() + ")"
}

func (s *sub) Trim() Evaluatable {
	// Drop a sub with zero
	leftTrim := s.left.Trim()
	rightTrim := s.right.Trim()

	dropLeft := false
	if leftTrim.IsConstant() {
		if leftTrim.Evaluate() == 0.0 {
			dropLeft = true
		}
	}
	dropRight := false
	if rightTrim.IsConstant() {
		if rightTrim.Evaluate() == 0.0 {
			dropRight = true
		}
	}

	if dropLeft && dropRight {
		return GetConstant(ConstantZero)
	} else if dropLeft {
		return rightTrim
	} else if dropRight {
		return leftTrim
	} else {
		return NodeSub(leftTrim, rightTrim)
	}
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
		return GetConstant(ConstantZero)
	}
}

func (m *multiply) String() string {
	return "(" + m.left.String() + " * " + m.right.String() + ")"
}

func (m *multiply) Trim() Evaluatable {
	// Kills a multiply with zero
	// Simplifies a multiply with one
	leftTrim := m.left.Trim()
	rightTrim := m.right.Trim()

	// Boolean flags:
	leftIsZero := false
	rightIsZero := false
	leftIsOne := false
	rightIsOne := false
	if leftTrim.IsConstant() {
		cValue := leftTrim.Evaluate()
		if cValue == 0.0 {
			leftIsZero = true
		} else if cValue == 1.0 {
			leftIsOne = true
		}
	}
	if rightTrim.IsConstant() {
		cValue := rightTrim.Evaluate()
		if cValue == 0.0 {
			rightIsZero = true
		} else if cValue == 1.0 {
			rightIsOne = true
		}
	}

	// Do the operations:
	if leftIsZero || rightIsZero {
		return GetConstant(ConstantZero)
	} else if leftIsOne {
		return rightTrim
	} else if rightIsOne {
		return leftTrim
	} else {
		return NodeMultiply(leftTrim, rightTrim)
	}
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
	if denominator == GetConstant(ConstantZero).Evaluate() {
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
				NodeMultiply(GetConstant(ConstantMinusOne), d.left),
				d.right.Diff(v),
			),
			NodeMultiply(d.right, d.right),
		)
	} else {
		return GetConstant(ConstantZero)
	}
}

func (d *divide) String() string {
	return "(" + d.left.String() + " / " + d.right.String() + ")"
}

func (d *divide) Trim() Evaluatable {
	// Kills a numerator equal to zero
	// Simplify a one denominator
	// Panic division by zero
	leftTrim := d.left.Trim()
	rightTrim := d.right.Trim()

	// Boolean flags:
	numIsZero := false
	denIsZero := false
	denIsOne := false
	if leftTrim.IsConstant() {
		cValue := leftTrim.Evaluate()
		if cValue == 0.0 {
			numIsZero = true
		}
	}
	if rightTrim.IsConstant() {
		cValue := rightTrim.Evaluate()
		if cValue == 0.0 {
			denIsZero = true
		} else if cValue == 1.0 {
			denIsOne = true
		}
	}

	// Do the operations:
	if denIsZero {
		panic("Expression contains division by zero!")
	} else if numIsZero {
		return GetConstant(ConstantZero)
	} else if denIsOne {
		return leftTrim
	} else {
		return NodeDivide(leftTrim, rightTrim)
	}
}
