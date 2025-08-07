package main

import "fmt"

type my_struct struct {
	name string
	age  int
}

func main() {
	var struct_value my_struct

	struct_value.age = 20
	struct_value.name = "Alan"

	fmt.Printf("Name: %s , Age: %v",struct_value.name,struct_value.age);
}