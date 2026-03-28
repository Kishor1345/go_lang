package main


import(
	"fmt"
)
//struct groups related feilds into one type

type User struct{
	ID int
	Name string
	Email string
	Age int
}

func main()  {
	u1 :=User{
		ID: 1,
		Name: "Kishore",
		Email: "kishoredinesh123@gmail.com",
		Age: 10,
	}
	fmt.Println(u1,u1.ID,u1.Email)

	//mutable by default
	u1.Age = 200
	fmt.Println(u1.Age)

	u2 :=User{
		Name: "kumar",
	}
	fmt.Println("partial user",u2)
}