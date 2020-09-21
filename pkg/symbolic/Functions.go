package symbolic

import (
	"math"
)

type pow struct {
	node
}

func NodePow(left, right Evaluatable) *pow {
	parent := node{left: left, right: right}
	return &pow{parent}
}

func (p *pow) Evaluate() float64 {
	base := p.left.Evaluate()
	exp := p.right.Evaluate()
	if base == 0.0 && exp == 0.0 {
		panic("Undetermined result 0^0")
	}
	return math.Pow(base, exp)
}

func (p *pow) Diff(v *Variable) Evaluatable {
	leftIsFunc := p.left.FunctionOf(v)
	rightIsFunc := p.right.FunctionOf(v)
	if leftIsFunc && rightIsFunc {
		// Full formula:
		return NodeAdd(
			NodeMultiply(
				NodeDivide(
					NodeMultiply(p.right, p),
					p.left,
				),
				p.left.Diff(v),
			),
			NodeMultiply(
				NodeMultiply(
					p,
					NodeLn(p.left),
				),
				p.right.Diff(v),
			),
		)
	} else if leftIsFunc {
		// Only base is func:
		return NodeMultiply(
			NodeMultiply(
				p.right,
				NodePow(
					p.left,
					NodeSub(p.right, GetConstant(CONSTANT_ONE)),
				),
			),
			p.left.Diff(v),
		)
	} else if rightIsFunc {
		// Only exp is func:
		return NodeMultiply(
			NodeMultiply(
				p,
				NodeLn(p.left),
			),
			p.right.Diff(v),
		)
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (p *pow) ToString() string {
	return "(" + p.left.ToString() + " ^ " + p.right.ToString() + ")"
}

type ln struct {
	node
}

func NodeLn(left Evaluatable) *ln {
	parent := node{left: left, right: nil}
	return &ln{parent}
}

func (l *ln) Evaluate() float64 {
	operand := l.left.Evaluate()
	if operand <= 0.0 {
		panic("Negative domain for Ln")
	}
	return math.Log(operand)
}

func (l *ln) Diff(v *Variable) Evaluatable {
	isFunc := l.left.FunctionOf(v)
	if isFunc {
		return NodeMultiply(
			NodeDivide(GetConstant(CONSTANT_ONE), l.left),
			l.left.Diff(v),
		)
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (l *ln) ToString() string {
	return "ln(" + l.left.ToString() + ")"
}

type sin struct {
	node
}

func NodeSin(left Evaluatable) *sin {
	parent := node{left: left, right: nil}
	return &sin{parent}
}

func (s *sin) Evaluate() float64 {
	return math.Sin(s.left.Evaluate())
}

func (s *sin) Diff(v *Variable) Evaluatable {
	isFunc := s.left.FunctionOf(v)
	if isFunc {
		return NodeMultiply(
			NodeCos(s.left),
			s.left.Diff(v),
		)
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (s *sin) ToString() string {
	return "sin(" + s.left.ToString() + ")"
}

type cos struct {
	node
}

func NodeCos(left Evaluatable) *cos {
	parent := node{left: left, right: nil}
	return &cos{parent}
}

func (c *cos) Evaluate() float64 {
	return math.Cos(c.left.Evaluate())
}

func (c *cos) Diff(v *Variable) Evaluatable {
	isFunc := c.left.FunctionOf(v)
	if isFunc {
		return NodeMultiply(
			NodeMultiply(
				GetConstant(CONSTANT_MINUS_ONE),
				NodeSin(c.left),
			),
			c.left.Diff(v),
		)
	} else {
		return GetConstant(CONSTANT_ZERO)
	}
}

func (c *cos) ToString() string {
	return "cos(" + c.left.ToString() + ")"
}
