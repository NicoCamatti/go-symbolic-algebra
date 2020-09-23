package test

import (
	"math"
	symb "symbolic-algebra/pkg/symbolic"
	"testing"
)

func TestNodeAdd(t *testing.T) {
	x := symb.CreateVariable("x")
	y := symb.CreateVariable("y")
	add := symb.NodeAdd(x, y)
	// 1.0 + 2.0 = 3.0
	x.SetValue(1.0)
	y.SetValue(2.0)
	got := add.Evaluate()
	if got != 3.0 {
		t.Error("Expected 3.0, got", got)
	}

	// 1.0 + (-1.0) = 0.0
	x.SetValue(1.0)
	y.SetValue(-1.0)
	got = add.Evaluate()
	if got != 0.0 {
		t.Error("Expected 0.0, got", got)
	}

	// 1.0 + (-4.0) = -3.0
	x.SetValue(1.0)
	y.SetValue(-4.0)
	got = add.Evaluate()
	if got != -3.0 {
		t.Error("Expected -3.0, got", got)
	}
}

func TestNodeSub(t *testing.T) {
	x := symb.CreateVariable("x")
	y := symb.CreateVariable("y")
	add := symb.NodeSub(x, y)
	// 1.0 - 2.0 = -1.0
	x.SetValue(1.0)
	y.SetValue(2.0)
	got := add.Evaluate()
	if got != -1.0 {
		t.Error("Expected -1.0, got", got)
	}

	// 1.0 - 1.0 = 0.0
	x.SetValue(1.0)
	y.SetValue(1.0)
	got = add.Evaluate()
	if got != 0.0 {
		t.Error("Expected 0.0, got", got)
	}

	// 1.0 - (-7.0) = 8.0
	x.SetValue(1.0)
	y.SetValue(-7.0)
	got = add.Evaluate()
	if got != 8.0 {
		t.Error("Expected 8.0, got", got)
	}
}

func TestNodeMultiply(t *testing.T) {
	x := symb.CreateVariable("x")
	y := symb.CreateVariable("y")
	mult := symb.NodeMultiply(x, y)
	// 2.5 * 4.0 = 10.0
	x.SetValue(2.5)
	y.SetValue(4.0)
	got := mult.Evaluate()
	if got != 10.0 {
		t.Error("Expected 10.0, got", got)
	}

	// 8.0 * (-1.0) = -8.0
	x.SetValue(8.0)
	y.SetValue(-1.0)
	got = mult.Evaluate()
	if got != -8.0 {
		t.Error("Expected -8.0, got", got)
	}

	// (-2.5) * (-8.0) = 20.0
	x.SetValue(-2.5)
	y.SetValue(-8.0)
	got = mult.Evaluate()
	if got != 20.0 {
		t.Error("Expected 20.0, got", got)
	}

	// (-2.5) * 0.0 = 0.0
	x.SetValue(-2.5)
	y.SetValue(0.0)
	got = mult.Evaluate()
	if got != 0.0 {
		t.Error("Expected 0.0, got", got)
	}
}

func TestNodeDivide(t *testing.T) {
	x := symb.CreateVariable("x")
	y := symb.CreateVariable("y")
	div := symb.NodeDivide(x, y)
	// 100.0 / 25.0 = 4.0
	x.SetValue(100.0)
	y.SetValue(25.0)
	got := div.Evaluate()
	if got != 4.0 {
		t.Error("Expected 4.0, got", got)
	}

	// 100 / (-25.0) = -4.0
	x.SetValue(100.0)
	y.SetValue(-25.0)
	got = div.Evaluate()
	if got != -4.0 {
		t.Error("Expected -4.0, got", got)
	}

	// Test division by zero:
	defer func() {
		if r := recover(); r == nil {
			t.Error("Division by zero did not panic")
		}
	}()
	x.SetValue(1.0)
	y.SetValue(0.0)
	div.Evaluate()
}

func TestCreateVariable(t *testing.T) {
	x := symb.CreateVariable("x")
	// Test name:
	got_name := x.GetName()
	if got_name != "x" {
		t.Error("Wrong name for variable x", got_name)
	}
	// Test default value:
	got_value := x.Evaluate()
	if got_value != 0.0 {
		t.Error("Expected default value 0.0, got", got_value)
	}
}

func TestSetValue(t *testing.T) {
	x := symb.CreateVariable("x")
	x.SetValue(11.67)
	got := x.Evaluate()
	if got != 11.67 {
		t.Error("Expected SetValue 11.67, got", got)
	}
}

func TestGetConstant(t *testing.T) {
	x := symb.GetConstant(symb.ConstantPi)
	got := x.Evaluate()
	if got != math.Pi {
		t.Error("Wrong value for Pi:", x.Evaluate())
	}
}

func TestGetConstantValue(t *testing.T) {
	x := symb.GetConstantValue(35.2)
	got_name := x.GetName()
	if got_name != "35.2" {
		t.Error("Wrong name for constant. Expected 35.2, got", got_name)
	}
	got_value := x.Evaluate()
	if got_value != 35.2 {
		t.Error("Wrong value for constant. Expected 35.2, got", got_value)
	}
}

func TestFunctionOf(t *testing.T) {
	x := symb.CreateVariable("x")
	y := symb.NodeAdd(
		symb.NodeMultiply(
			symb.GetConstantValue(2.0),
			x,
		),
		symb.GetConstantValue(3.0),
	)
	got := y.FunctionOf(x)
	if !got {
		t.Error("y should be function of x")
	}
}

func TestFunctionOfWithConstantSameName(t *testing.T) {
	x := symb.CreateVariable("x")
	e := symb.CreateVariable("e")
	y := symb.NodeAdd(
		symb.NodeMultiply(
			symb.GetConstantValue(2.0),
			x,
		),
		// Conflicts with variable "z", but code should know the difference
		symb.GetConstant(symb.ConstantE),
	)
	got := y.FunctionOf(e)
	if got {
		t.Error("y should not be function of variable e")
	}
}

func TestDiffWithAdd(t *testing.T) {
	x := symb.CreateVariable("x")
	// y = x + 1
	y := symb.NodeAdd(
		x,
		symb.GetConstant(symb.ConstantOne),
	)
	expr := y.Diff(x).String()
	if expr != "1" {
		t.Error("Expected 1 for derivative, got", expr)
	}
}

func TestDiffWithSub(t *testing.T) {
	x := symb.CreateVariable("x")
	// y = 1 - x
	y := symb.NodeSub(
		symb.GetConstant(symb.ConstantOne),
		x,
	)
	expr := y.Diff(x).String()
	if expr != "(-1 * 1)" {
		t.Error("Expected (-1 * 1) for derivative, got", expr)
	}
}

func TestDiffWithMultiply(t *testing.T) {
	x := symb.CreateVariable("x")
	// y = pi * x
	y := symb.NodeMultiply(
		symb.GetConstant(symb.ConstantPi),
		x,
	)
	expr := y.Diff(x).String()
	if expr != "(pi * 1)" {
		t.Error("Expected (pi * 1) for derivative, got", expr)
	}
}

func TestDiffWithDivide(t *testing.T) {
	x := symb.CreateVariable("x")
	// y = 1 / x
	y := symb.NodeDivide(
		symb.GetConstant(symb.ConstantOne),
		x,
	)
	expr := y.Diff(x).String()
	if expr != "(((-1 * 1) * 1) / (x * x))" {
		t.Error("Expected (((-1 * 1) * 1) / (x * x)) for derivative, got", expr)
	}
}

func TestNodePow(t *testing.T) {
	base := symb.CreateVariable("x")
	exp := symb.CreateVariable("n")
	pow := symb.NodePow(base, exp)

	// 2.0 ^ 10.0 = 1024.0
	base.SetValue(2.0)
	exp.SetValue(10.0)
	got := pow.Evaluate()
	if got != 1024.0 {
		t.Error("Expected 1024.0, got", got)
	}

	// 10.0 ^ 0.0 = 1.0
	base.SetValue(10.0)
	exp.SetValue(0.0)
	got = pow.Evaluate()
	if got != 1.0 {
		t.Error("Expected 1.0, got", got)
	}

	// 1024.0 ^ 0.5 = 32.0
	base.SetValue(1024.0)
	exp.SetValue(0.5)
	got = pow.Evaluate()
	if got != 32.0 {
		t.Error("Expected 32.0, got", got)
	}

	// 0.5 ^ -1.0 = 2.0
	base.SetValue(0.5)
	exp.SetValue(-1.0)
	got = pow.Evaluate()
	if got != 2.0 {
		t.Error("Expected 2.0, got", got)
	}
}

func TestNodeLn(t *testing.T) {
	x := symb.CreateVariable("x")
	ln := symb.NodeLn(x)

	// ln(1.0) = 0.0
	x.SetValue(1.0)
	got := ln.Evaluate()
	if got != 0.0 {
		t.Error("Expected 0.0, got", got)
	}

	// ln(e) = 1.0
	x.SetValue(math.E)
	got = ln.Evaluate()
	if got != 1.0 {
		t.Error("Expected 1.0, got", got)
	}

	// panic if zero:
	defer func() {
		if r := recover(); r == nil {
			t.Error("Ln should have panic")
		}
	}()
	x.SetValue(0.0)
	ln.Evaluate()

	// panic if less than zero:
	x.SetValue(-1.0)
	ln.Evaluate()
}

func TestNodeSin(t *testing.T) {
	x := symb.CreateVariable("x")
	sin := symb.NodeSin(x)

	// sin(0) = 0.0
	x.SetValue(0.0)
	got := sin.Evaluate()
	if math.Abs(got-0.0) > 1e-10 {
		t.Error("Expected 0.0, got", got)
	}

	// sin(2*pi) = 0.0
	x.SetValue(math.Pi * 2.0)
	got = sin.Evaluate()
	if math.Abs(got-0.0) > 1e-10 {
		t.Error("Expected 0.0, got", got)
	}

	// sin(-2*pi) = 0.0
	x.SetValue(-math.Pi * 2.0)
	got = sin.Evaluate()
	if math.Abs(got-0.0) > 1e-10 {
		t.Error("Expected 0.0, got", got)
	}
}
func TestNodeCos(t *testing.T) {
	x := symb.CreateVariable("x")
	cos := symb.NodeCos(x)

	// cos(0) = 1.0
	x.SetValue(0.0)
	got := cos.Evaluate()
	if math.Abs(got-1.0) > 1e-10 {
		t.Error("Expected 0.0, got", got)
	}

	// cos(2*pi) = 1.0
	x.SetValue(math.Pi * 2.0)
	got = cos.Evaluate()
	if math.Abs(got-1.0) > 1e-10 {
		t.Error("Expected 0.0, got", got)
	}

	// cos(-2*pi) = 1.0
	x.SetValue(-math.Pi * 2.0)
	got = cos.Evaluate()
	if math.Abs(got-1.0) > 1e-10 {
		t.Error("Expected 0.0, got", got)
	}
}

func TestDiffWithPow(t *testing.T) {
	x := symb.CreateVariable("x")
	two := symb.GetConstantValue(2.0)

	// d(x^2)/dx = 2x:
	x2 := symb.NodePow(x, two)
	expr := x2.Diff(x).String()
	if expr != "((2 * (x ^ (2 - 1))) * 1)" {
		t.Error("Expected ((2 * (x ^ (2 - 1))) * 1), got", expr)
	}

	// d(x^x)/dx = full formula:
	xx := symb.NodePow(x, x)
	expr = xx.Diff(x).String()
	if expr != "((((x * (x ^ x)) / x) * 1) + (((x ^ x) * ln(x)) * 1))" {
		t.Error("Expected ((((x * (x ^ x)) / x) * 1) + (((x ^ x) * ln(x)) * 1)), got", expr)
	}
}

func TestDiffWithLn(t *testing.T) {
	x := symb.CreateVariable("x")
	two := symb.GetConstantValue(2.0)
	ln := symb.NodeLn(
		symb.NodeMultiply(two, x),
	)

	expr := ln.Diff(x).String()
	if expr != "((1 / (2 * x)) * (2 * 1))" {
		t.Error("Expected ((1 / (2 * x)) * (2 * 1)), got", expr)
	}
}

func TestDiffWithSin(t *testing.T) {
	x := symb.CreateVariable("x")
	two := symb.GetConstantValue(2.0)
	sin := symb.NodeSin(
		symb.NodeMultiply(two, x),
	)

	expr := sin.Diff(x).String()
	if expr != "(cos((2 * x)) * (2 * 1))" {
		t.Error("Expected (cos((2 * x)) * (2 * 1)), got", expr)
	}
}

func TestDiffWithCos(t *testing.T) {
	x := symb.CreateVariable("x")
	two := symb.GetConstantValue(2.0)
	cos := symb.NodeCos(
		symb.NodeMultiply(two, x),
	)

	expr := cos.Diff(x).String()
	if expr != "((-1 * sin((2 * x))) * (2 * 1))" {
		t.Error("Expected ((-1 * sin((2 * x))) * (2 * 1)), got", expr)
	}
}
