package main

import ("fmt"
				pkg "go_tutorials/tutorial_9/pkg")

func main() {
	var c = make(chan int)
	go goprocess(c)
	for i := range c {
		fmt.Println(i);
	}

	pkg.Result();
}

func goprocess(c chan int) {
	defer close(c)
	for i := 0;i < 5; i++ {
		c <- i
	}
}