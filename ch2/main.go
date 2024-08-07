package main

import (
	"fmt"
	// "math/cmplx"
)

func main() {
	// x := complex(2.5, 3.1)
	// y := complex(10.2, 2)
	// // fmt.Println(x + y)
	// // fmt.Println(x - y)
	// // fmt.Println(x * y)
	// // fmt.Println(x / y)
	// // fmt.Println(real(x))
	// // fmt.Println(imag(x))
	// // fmt.Println(cmplx.Abs(x))
	// var a int = 10
	// var b float64 = 30.2
	// var sum1 float64 = float64(a) + b
	// var sum2 int = a + int(b)
	// // fmt.Println(sum1,sum2)

	i := 20
	var f float64 = float64(i)
	fmt.Println(i)
	fmt.Println(f)
	const value = 10
	i = value
	f = float64(value)
	fmt.Println(i)
	fmt.Println(f)

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b+1
	smallI = smallI + 1
	bigI = bigI + 1
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

}
