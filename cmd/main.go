package main

import (
	"fmt"
	symb "symbolic-algebra/pkg/symbolic"
)

func main() {
	fmt.Println("Starting test:")

	x := symb.CreateVariable("x")
	zero := symb.GetConstant(symb.ConstantZero)
	sum := symb.NodeAdd(x, zero)
	supersum := symb.NodeAdd(sum, sum)

	fmt.Println("Original:", supersum)
	fmt.Println("Original is function of X", supersum.FunctionOf(x))
	fmt.Println("Original is constant", supersum.IsConstant())
	fmt.Println("Trimmed:", supersum.Trim())
	fmt.Println("Trimmed is constant:", supersum.Trim().IsConstant())

	fmt.Println("X is constant?", x.IsConstant())
	fmt.Println("0 is constant?", zero.IsConstant())

}
