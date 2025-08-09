package main

import (
	"fmt"
	datapkg "go_tutorials/tutorial_5/data_pkg"
	mypkg2 "go_tutorials/tutorial_5/my_pkg2"
)

type my_struct struct {
	name string
	age  int
}

func (p my_struct) display() {
	fmt.Printf("This is a function attached to the struct my_struct:\n name: %s, age: %v\n", p.name, p.age)
}

func (p *my_struct) add() {
	p.age++
}

func main() {
	var struct_value my_struct

	struct_value.age = 20
	struct_value.name = "Alan"

	fmt.Printf("Name: %s , Age: %v\n", struct_value.name, struct_value.age)
	struct_value.display()
	struct_value.add()
	struct_value.display()

	// Create a struct for the Loop function
	var data_struct datapkg.NewStruct
	data_struct.Name = "John"
	data_struct.Age = 25
	data_struct.Id = 1

	mypkg2.Loop(data_struct)
}
