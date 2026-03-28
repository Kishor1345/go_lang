package main

import (
	 "go-modules/internal/greet"
	 "fmt"
)
func main()  {
	msg1 :=greet.Hello("kishore")
	fmt.Println(msg1)
}