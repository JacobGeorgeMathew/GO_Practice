package main

import "fmt"

func main() {
	// var num int = 100
	// var num2 float32 = 123.321

	// fmt.Println("Hello World ", num)
	// fmt.Println(num2)
	// var result float32 = num2 + float32(num)

	// fmt.Println(result);

	// var myString string = "Hello \n Jacob";
	// fmt.Println(myString);
	// fmt.Println(len("ABC"));

	// myVar := "text";
	// fmt.Println(myVar);
	display("Hi");
	var sum, diff, err = Add(120,50);
	if err != nil {
		fmt.Printf("Error occured \n");
	}else {
	fmt.Printf("Sum %v \n difference %v",sum,diff);
	}	
}

func display(para string) {
	fmt.Println("Hello World",para);
}

func Add(a int , b int) (int, int, error) {
	var err error;
	if a == 120 {
		return 0 , 0, err;
	}
	var result int = a + b;
	var diff int = a - b;
	return result,diff,err;
}
