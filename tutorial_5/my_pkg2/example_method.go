package mypkg2

import (
	"fmt"
	datapkg "go_tutorials/tutorial_5/data_pkg"
)

func Loop(p datapkg.NewStruct) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Lap %v, name: %s, age: %v\n", i, p.Name, p.Age)
	}
	datapkg.Model();
}
