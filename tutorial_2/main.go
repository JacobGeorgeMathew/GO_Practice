package main

import "fmt"

func main() {
//  var Arr1 [3]int32 = [3]int32{1,2,3};
//  Arr2 := [3]int32{4,5,6};
//  fmt.Println((Arr1),(Arr2)); 
//  fmt.Println(Arr1[0:2]);

 var s []int
fmt.Println(len(s), cap(s)) // 0 0

s = append(s, 1)
fmt.Println(s, len(s), cap(s)) // [1] 1 1

s = append(s, 2, 3, 4)
fmt.Println(s, len(s), cap(s)) // [1 2 3 4] 4 4

s = append(s, 5)
fmt.Println(s, len(s), cap(s)) // [1 2 3 4 5] 5 8 (capacity doubled)

}