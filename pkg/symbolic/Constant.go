package symbolic

import (
	"math"
)

const (
	CONSTANT_ZERO      string = "0"
	CONSTANT_ONE       string = "1"
	CONSTANT_MINUS_ONE string = "-1"
	CONSTANT_E         string = "e"
	CONSTANT_PI        string = "pi"
)

// Gets a constant from the given input string. Only supports the already defined const strings, otherwise panic
func GetConstant(c string) *Constant {
	var rConst Constant
	switch c {
	case CONSTANT_ZERO:
		rConst = Constant{name: c, value: 0.0}
	case CONSTANT_ONE:
		rConst = Constant{name: c, value: 1.0}
	case CONSTANT_MINUS_ONE:
		rConst = Constant{name: c, value: -1.0}
	case CONSTANT_E:
		rConst = Constant{name: c, value: math.E}
	case CONSTANT_PI:
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
	return GetConstant(CONSTANT_ZERO)
}

func (c *Constant) ToString() string {
	return c.name
}
