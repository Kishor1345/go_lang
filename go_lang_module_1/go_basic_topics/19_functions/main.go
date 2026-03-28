package main

import(
	"fmt"
)

//single return value
func add(a int,b int) int{
	return a+b
}

//double return
func SumandProduct(a int,b int)(int,int){
	sum :=a+b
	product :=a* b
	return sum ,product
}

func main()  {
	sum1 := add(10,20)
	fmt.Println(sum1)

	//ignore by using underscore
	sum,_ := SumandProduct(10,20)
	fmt.Println(sum)


}