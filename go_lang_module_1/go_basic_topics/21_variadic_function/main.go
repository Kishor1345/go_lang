package main

import(
	"fmt"
)

//variadic function pass n number of paramter fo function
func sumAll(nums ...int)int{
	total := 0
	for _,n :=range nums{
		total +=n
	}
	return total
}

func main()  {
	s:=sumAll(1,2,3,4,5)
	fmt.Println(s)

	values :=[]int{10,20,30}
	fmt.Println(sumAll(values...))


	//anonymous function
	result := func (n int)int  {
		return n*2
	}
	fmt.Println(result(2))

	//IIFE

	result1 :=func (a,b int) int {
		return a+b
	}(5,10)

	fmt.Println(result1)

	result2 :=map[string]int{
		"kishore":22,
		"dinesh":23,
	}

	if _,ok :=result2["kishore"];!ok{
		fmt.Println("error")
	}else{
		fmt.Println("no error")
	}

}