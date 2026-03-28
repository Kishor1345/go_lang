package main


import (
	"fmt"
)

type User struct{
	Name string
	Age int
}

func main()  {
	u:=User{
		Name: "Kishore",
		Age: 22,
	}	
	fmt.Println(u.Age)
	u.Birthday()
	fmt.Println(u.Age)
}

func (u *User) Birthday() {
	u.Age++
}