package main

import ("fmt"
				"go_tutorials/tutorial_3/mypkg");

func main() {
  var my_string string = "Hello";

	for i , v := range my_string{
	fmt.Println(i , v);
	}

	s := "Hello, 世界"
r := []rune(s)

fmt.Println(len(s))     // Length in bytes
fmt.Println(len(r))     // Length in runes (characters)

for i, r := range "Go语言" {
    fmt.Printf("%d: %c\n", i, r)
}
mypkg.Malayalam();
}