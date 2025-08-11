package pkg

import "fmt"

func Sum(nums []int, ch chan int) {
    total := 0
    for _, n := range nums {
        total += n
    }
    ch <- total // send result to channel
}

func Result() {
    Nums := []int{1, 2, 3, 4}
    ch := make(chan int)

    go Sum(Nums[:2], ch)
    go Sum(Nums[2:], ch)

    x, y := <-ch, <-ch
		fmt.Println("x:",x)
		fmt.Println("y:",y)
    fmt.Println("Total:", x+y)
}
