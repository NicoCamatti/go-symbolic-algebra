package symbolic

// A Variable is a container for a not constant float64 number
type Variable struct {
	Constant
}

func CreateVariable(name string) *Variable {
	parent := Constant{name: name, value: 0.0}
	return &Variable{parent}
}

// SetValue of this Variable with the given float64 argument
func (v *Variable) SetValue(newValue float64) {
	v.value = newValue
}

// Returns true if this variable matches the one being checked
func (this *Variable) FunctionOf(v *Variable) bool {
	return this.name == (*v).name
}

// Returns false because it is variable
func (v *Variable) IsConstant() bool {
	return false
}

// Returns 1.0 if FunctionOf is true 0.0 otherwise
func (this *Variable) Diff(v *Variable) Evaluatable {
	if this.FunctionOf(v) {
		return GetConstant(ConstantOne)
	} else {
		return GetConstant(ConstantZero)
	}
}

func (v *Variable) Trim() Evaluatable {
	return v
}
