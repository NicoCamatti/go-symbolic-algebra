package symbolic

import (
	"fmt"
)

// A Evaluatable interface for nodes in a binary expression tree
type Evaluatable interface {
	fmt.Stringer
	Evaluate() float64
	FunctionOf(v *Variable) bool
	Diff(v *Variable) Evaluatable
}
