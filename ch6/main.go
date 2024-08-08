package main

import "fmt"


type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{FirstName: firstName, LastName: lastName, Age: age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age}
}

func main() {

	p1 := MakePerson("John", "Doe", 30)
	p2 := MakePersonPointer("Jane", "Doe", 25)

	fmt.Println(p1)
	fmt.Println(p2)
}