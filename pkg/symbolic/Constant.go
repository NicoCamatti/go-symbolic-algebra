package symbolic

import (
	"math"
)

const (
	ConstantZero     string = "0"
	ConstantOne      string = "1"
	ConstantMinusOne string = "-1"
	ConstantE        string = "e"
	ConstantPi       string = "pi"
)

// Gets a constant from the given input string. Only supports the already defined const strings, otherwise panic
func GetConstant(c string) *Constant {
	var rConst Constant
	switch c {
	case ConstantZero:
		rConst = Constant{name: c, value: 0.0}
	case ConstantOne:
		rConst = Constant{name: c, value: 1.0}
	case ConstantMinusOne:
		rConst = Constant{name: c, value: -1.0}
	case ConstantE:
		rConst = Constant{name: c, value: math.E}
	case ConstantPi:
		rConst = Constant{name: c, value: math.Pi}
	default:
		panic("Unknown constant")
	}
	return &rConst
}

// Get a new constant from the given name and value
func GetCustomConstant(c string, v float64) *Constant {
	return &Constant{name: c, value: v}
}

// A Constant value float64
type Constant struct {
	name  string
	value float64
}

func (c *Constant) GetName() string {
	return c.name
}

func (c *Constant) Evaluate() float64 {
	return c.value
}

// Automatically returns false if got to constant leaf node
func (c *Constant) FunctionOf(v *Variable) bool {
	return false
}

// Automatically returns 0.0 if got to constant leaf node
func (c *Constant) Diff(v *Variable) Evaluatable {
	return GetConstant(ConstantZero)
}

func (c *Constant) String() string {
	return c.name
}
