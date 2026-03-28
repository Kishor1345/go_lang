package main

import(
	"fmt"
)
func main()  {
	//array is fixed size
	//slice is dynamic size

	//fixed size can not grow
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 30

	fmt.Println(marks)

	//array literal
	res :=[5]int{1,2,3,4,5}
	fmt.Println(len(res))
}