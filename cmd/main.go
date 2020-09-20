package main

import (
	"fmt"
	symb "symbolic-algebra/pkg/symbolic"
)

func main() {
	fmt.Println("Starting test:")

	// y = 4x -1
	x := symb.CreateVariable("x")
	y := symb.NodeSub(symb.NodeMultiply(symb.GetCustomConstant("4", 4.0), x), symb.GetCustomConstant("1", 1.0))
	fmt.Println("y:", y.ToString())
	fmt.Println("X", "Y")
	for i := 0; i < 10; i++ {
		x.SetValue(float64(i))
		fmt.Println(x.Evaluate(), y.Evaluate())
	}

	// y = 2x^2 - x + 3
	y2 := symb.NodeAdd(
		symb.NodeSub(
			symb.NodeMultiply(
				symb.GetCustomConstant("2", 2.0),
				symb.NodeMultiply(x, x),
			),
			x,
		),
		symb.GetCustomConstant("3", 3.0),
	)
	fmt.Println("y2:", y2.ToString())
	// fmt.Println("X", "Y")
	// for i := 0; i < 10; i++ {
	// 	x.SetValue(float64(i))
	// 	fmt.Println(x.Evaluate(), y2.Evaluate())
	// }

	// Testing derivative:
	y2_ := y2.Diff(x)
	fmt.Println("y2_", y2_.ToString())
	fmt.Println("X", "Y")
	for i := 0; i < 10; i++ {
		x.SetValue(float64(i))
		fmt.Println(x.Evaluate(), y2_.Evaluate())
	}

	// Testing second derivative:
	y2__ := y2.Diff(x).Diff(x)
	fmt.Println("y2__", y2__.ToString())
	fmt.Println("X", "Y")
	for i := 0; i < 10; i++ {
		x.SetValue(float64(i))
		fmt.Println(x.Evaluate(), y2__.Evaluate())
	}

}
