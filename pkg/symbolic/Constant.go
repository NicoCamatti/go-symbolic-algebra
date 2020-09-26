package symbolic

import (
	"fmt"
	"math"
)

// A constant pool of exported string names that refer to a constant numeric
const (
	ConstantZero     string = "0"
	ConstantOne      string = "1"
	ConstantMinusOne string = "-1"
	ConstantE        string = "e"
	ConstantPi       string = "pi"
)

// A pool of Constants
var constantPool = map[string]Constant{
	ConstantZero:     {ConstantZero, 0.0},
	ConstantOne:      {ConstantOne, 1.0},
	ConstantMinusOne: {ConstantMinusOne, -1.0},
	ConstantE:        {ConstantE, math.E},
	ConstantPi:       {ConstantPi, math.Pi},
}

// Gets a constant from the given input string. Only supports the already defined const strings, otherwise panic
func GetConstant(c string) *Constant {
	var rConst Constant
	rConst, ok := constantPool[c]
	if ok {
		return &rConst
	}
	panic("Unknown constant")
}

// Get a constant from the given value
func GetConstantValue(v float64) (rConst *Constant) {
	name := fmt.Sprintf("%g", v) // Creates a name for this constant v
	defer func() {
		if r := recover(); r != nil {
			// If panic add the new constant to the pool and return it:
			rConst = &Constant{name: name, value: v}
			constantPool[name] = *rConst
		}
	}()
	rConst = GetConstant(name) // May panic if not mapped constant
	return
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

// Automatically returns true when it is constant
func (c *Constant) IsConstant() bool {
	return true
}

// Automatically returns 0.0 if got to constant leaf node
func (c *Constant) Diff(v *Variable) Evaluatable {
	return GetConstant(ConstantZero)
}

func (c *Constant) String() string {
	return c.name
}

func (c *Constant) Trim() Evaluatable {
	return c
}
