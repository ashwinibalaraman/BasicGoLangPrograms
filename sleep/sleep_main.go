package main

import (
	"fmt"
	"time"
)
import s "github.com/assignment1/sleep/sleepTime"

func main() {
	fmt.Println("hello at ", time.Now())
	s.Sleep(time.Second * 5)
	fmt.Println("world after 5 seconds sleep", time.Now())
}
