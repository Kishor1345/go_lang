package main

import(
	"fmt"
)
func main()  {
	//map[keytype]value type

	ages :=map[string]int{
		"kishore":22,
		"kumar":43,
	}
	fmt.Println(ages["kishore"],len(ages))

	//make(map[k]v)

	var scores map[string]int //nil map
	fmt.Println(scores)

	scores = make(map[string]int)

	scores["math"]=90
	fmt.Println(scores)

	users :=map[string]string{
		"u1":"kishore",
		"u2":"kumar",
	}
	fmt.Println(users)

	delete(users,"u2")//delete key and value in map
	fmt.Println(users)


	
}