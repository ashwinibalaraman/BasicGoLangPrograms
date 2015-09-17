package main

import "fmt"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("recovered from panic")
		}

	}()

	s := []int{1, 2, 3}
	pos := 4
	if pos > len(s) {
		panic("Array out of bounds")

	}

	fmt.Println(s[pos])

}
