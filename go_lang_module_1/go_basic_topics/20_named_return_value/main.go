package main


import(
	"fmt"
)

func divide(a int,b int)(q int ,r int){
	q = a/b
	r = a%b

	return 
}
func main()  {
	q,r:=divide(10,20)
	fmt.Println(q,r)
}