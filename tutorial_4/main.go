package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("First loop",i);
	}

	var array = [5]string{"a","b","c","d","e"};

	for i , v := range array {
		fmt.Printf("Second loop %v\t%v\n",i,v);
	}
	
}