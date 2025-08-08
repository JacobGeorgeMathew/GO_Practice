package mypkg2

import ("fmt" 
				"go_tutorials/tutorial_5/data_pkg")

func Loop(p *data_pkg.new_struct) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Lap %v, name: %s, age: %v",i,p.name,p.age);
	}
}