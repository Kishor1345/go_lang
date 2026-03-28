package main

import(
	"fmt"
)

func main()  {
	isLogged := true
	isAdmin := false
	hasAssign := true

	//And &&
	canOpenDashBoard := isLogged && hasAssign

	canDeletePost := isAdmin || (isLogged && hasAssign)

	fmt.Println(canOpenDashBoard,canDeletePost)

	age :=25
	isAdult := age>15
	fmt.Println(isAdult)
}