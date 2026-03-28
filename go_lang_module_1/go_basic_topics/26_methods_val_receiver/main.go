package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{
		Name: "kishore",
		Age:  22,
	}
	fmt.Println(u.Intro())
}

// val receiver means this method receive a copy of the user
func (u User) Intro() string {
	return fmt.Sprintf("Hi Iam %s",u.Name)
}