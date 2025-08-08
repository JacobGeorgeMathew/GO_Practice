package datapkg

type new_struct struct {
	name string
	age int
	id int
}

func Model(){
	var example new_struct;
	example.id = 10
	example.age = 20
	example.name = "default"
}