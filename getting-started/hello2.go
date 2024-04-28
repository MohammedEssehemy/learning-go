package main

import (
	"example/hello/hello3"
	"fmt"

	"rsc.io/quote"
)

func greet() {
	hello3.Greet3()
	fmt.Println(quote.Go())
	fmt.Println("Hello, Greet!")
}
