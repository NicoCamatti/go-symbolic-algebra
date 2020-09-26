package symbolic

import (
	"math"
)

type pow struct {
	node
}

func NodePow(left, right Evaluatable) Evaluatable {
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
					NodeSub(p.right, GetConstant(ConstantOne)),
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
		return GetConstant(ConstantZero)
	}
}

func (p *pow) String() string {
	return "(" + p.left.String() + " ^ " + p.right.String() + ")"
}

func (p *pow) Trim() Evaluatable {
	return nil // TODO implement
}

type ln struct {
	node
}

func NodeLn(left Evaluatable) Evaluatable {
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
			NodeDivide(GetConstant(ConstantOne), l.left),
			l.left.Diff(v),
		)
	} else {
		return GetConstant(ConstantZero)
	}
}

func (l *ln) String() string {
	return "ln(" + l.left.String() + ")"
}

func (l *ln) Trim() Evaluatable {
	return nil // TODO implement
}

type sin struct {
	node
}

func NodeSin(left Evaluatable) Evaluatable {
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
		return GetConstant(ConstantZero)
	}
}

func (s *sin) String() string {
	return "sin(" + s.left.String() + ")"
}

func (s *sin) Trim() Evaluatable {
	return nil // TODO implement
}

type cos struct {
	node
}

func NodeCos(left Evaluatable) Evaluatable {
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
				GetConstant(ConstantMinusOne),
				NodeSin(c.left),
			),
			c.left.Diff(v),
		)
	} else {
		return GetConstant(ConstantZero)
	}
}

func (c *cos) String() string {
	return "cos(" + c.left.String() + ")"
}

func (c *cos) Trim() Evaluatable {
	return nil // TODO implement
}
