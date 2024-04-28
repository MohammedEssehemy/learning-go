package main

import "fmt"

func sum(s []int, c chan int) {
	fmt.Println(s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(s[:len(s)/2], "before first")
	go sum(s[:len(s)/2], c)
	fmt.Println(s[len(s)/2:], "after first")
	go sum(s[len(s)/2:], c)
	fmt.Println("after second")
	x, y := <-c, <-c // receive from c
	fmt.Println("after all")
	fmt.Println(x, y, x+y)
}
