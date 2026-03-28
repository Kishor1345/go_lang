package main

import (
	"fmt"
	"strings"
)

func main()  {
	
	firstName := "kishore"
	lastName := "K"
	fullName := firstName +" "+ lastName
	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}