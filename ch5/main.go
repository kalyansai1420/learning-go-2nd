package main

import (
	// "errors"
	"fmt"
	"strconv"
	// "os"
)

// func addTo(base int,vals ...int) []int{
// 	out := make([]int,0,len(vals))
// 	for _,v := range vals {
// 		out = append(out,base+v)
// 	}
// 	return out
// }

// func divAndRemainder(num,denom int) (int,int,error) {
// 	if denom == 0 {
// 		return 0, 0, errors.New("cannot divide by zero")
// 	}
// 	return num/denom, num%denom,nil
// }

// func f1(a string) int {
// 	return len(a)
// }

// func f2(a string) int {
// 	total := 0
// 	for _,v := range a {
// 		total +=int(v)
// 	}
// 	return total
// }

func add(i int, j int ) int { return i + j }
func sub(i int, j int ) int { return i - j }
func mul(i int, j int ) int { return i * j }
func div(i int, j int ) int { return i / j }


var opMap = map[string]func(int,int)int{
	"+":add,
	"-":sub,
	"*":mul,
	"/":div,

}

func main() {

	expressions := [][]string{
		{"2","+","3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2","/","3"},
		{"2","%","3"},
		{"two","+","three"},
		{"5"},
	}

	for _,expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ",expression)
			continue
		}

		p1,err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op:=expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator: ",op)
			continue
		}

		p2,err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1,p2)
		fmt.Println(result)




	}

	// fmt.Println(addTo(3))
	// fmt.Println(addTo(3,2))
	// fmt.Println(addTo(3,2,4,6,8))
	// a := []int{4,3}
	// fmt.Println(addTo(3,a...))
	// fmt.Println(addTo(3,[]int{1,2,3,4,5}...))

	// result, remainder, err := divAndRemainder(5,2)
	// if err!= nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(result,remainder)

	// var myFuncVariable func(string) int
	// myFuncVariable = f1
	// result := myFuncVariable("Hello")
	// fmt.Println(result)

	// myFuncVariable = f2
	// result = myFuncVariable("Hello")
	// fmt.Println(result)

}