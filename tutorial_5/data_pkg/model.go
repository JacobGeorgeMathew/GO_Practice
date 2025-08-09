package datapkg

import "fmt"

type NewStruct struct {
	Name string
	Age  int
	Id   int
}

func Model() {
	var example NewStruct
	example.Id = 10
	example.Age = 20
	example.Name = "default"
	fmt.Printf("name: %s,Age: %v,Id: %v",example.Name,example.Age,example.Id);
}
