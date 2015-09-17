package main

import (
	"flag"
	"fmt"
)

import f "github.com/assignment1/fibo/fibonacci"

func main() {
	var n int
	flag.IntVar(&n, "n", 5, "an int")
	flag.Parse()

	for i := 0; i < n; i++ {
		fmt.Println(f.Fib(i))
	}

}
