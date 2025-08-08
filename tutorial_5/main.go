package main

import "fmt"

type my_struct struct {
	name string
	age  int
}

func (p my_struct) display() {
	fmt.Printf("This is a function attached to the struct my_struct:\n name: %s, age: %v\n",p.name,p.age);
}

func (p *my_struct) add() {
	p.age++;
}

func main() {
	var struct_value my_struct

	struct_value.age = 20
	struct_value.name = "Alan"

	fmt.Printf("Name: %s , Age: %v\n",struct_value.name,struct_value.age);
	struct_value.display();
	struct_value.add();
	struct_value.display();
}