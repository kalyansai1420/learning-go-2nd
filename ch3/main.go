package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	id        int
}



func main() {
	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x,10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x,20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x,30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x,40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x,50)
	// fmt.Println(x, len(x), cap(x))

	// data := []int{2, 4, 6, 8}
	// fmt.Println(data)

	// x := []string{"a","b","c","d"}
	// y := x[:2]
	// z := x[1:]
	// d := x[1:3]
	// e := x[:]
	// fmt.Println("x: ",x)
	// fmt.Println("y: ",y)
	// fmt.Println("z: ",z)
	// fmt.Println("d: ",d)
	// fmt.Println("e: ",e)

	// x :=  []string{"a","b","c","d"}
	// y := x[:2]
	// z := x[1:]

	// x[1] = "y"
	// y[0] = "x"
	// z[1] = "z"

	// fmt.Println("x: ",x)
	// fmt.Println("y: ",y)
	// fmt.Println("z: ",z)

	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subSlice1 := greetings[:2]
	subSlice2 := greetings[1:4]
	subSlice3 := greetings[3:]

	fmt.Println("greetings: ",greetings)
	fmt.Println("subSlice1: ",subSlice1)
	fmt.Println("subSlice2: ",subSlice2)
	fmt.Println("subSlice3: ",subSlice3)

	message := "Hello 世界 and दुनिया"
	fmt.Printf("The fourth rune is: %c\n", []rune(message)[3])
	
	employee1 := Employee{"John", "Doe" ,1}
	employee2 := Employee{firstName: "Jane", lastName:"Smith",id:2}

	var employee3 Employee
	employee3.firstName = "Denis"
	employee3.lastName = "Rich"
	employee3.id = 3

	fmt.Println("Employee 1: ",employee1)
	fmt.Println("Employee 2: ",employee2)
	fmt.Println("Employee 3: ",employee3)

}

