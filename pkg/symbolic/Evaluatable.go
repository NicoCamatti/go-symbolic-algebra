package symbolic

// A Evaluatable method to get a float64 value
type Evaluatable interface {
	Evaluate() float64
	FunctionOf(v *Variable) bool
	Diff(v *Variable) Evaluatable
	ToString() string
}
