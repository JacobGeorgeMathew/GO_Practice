package main

import "fmt"

func main() {
	var data = map[string]int{"abc": 10, "def": 20, "ghi": 30, "jkl": 40}

	for i, v := range data {
		fmt.Printf("key: %s , value: %v\n",i,v);
	}

	data["mno"] = 50

	val1 , exist1 := data["mno"];

	if exist1 {
		fmt.Printf("exist: %v\n",val1)
	}  else   
	 {
		fmt.Printf("does not exist:\n")
	}

	delete(data,"mno");

	val2 , exist2 := data["mno"];

	if exist2 {
		fmt.Printf("exist: %v\n",val2)
	}else {
		fmt.Printf("does not exist:\n")
	}
}