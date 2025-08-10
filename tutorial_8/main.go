package main

import ("fmt"
				pointerpkg "go_tutorials/tutorial_8/pointer_pkg")

func main() {
	var p *int32 = new(int32)
	var k *int32 = new(int32)
	var i int32 = 100
	fmt.Printf("p: Pointer: %v, Value: %v\n",p,*p);
	*p = 500;
	fmt.Printf("p: Pointer: %v, Value: %v\n",p,*p);
	k = &i
	fmt.Printf("k: Pointer: %v, Value: %v\n",k,*k);

	pointerpkg.Pointer_Practice();
}