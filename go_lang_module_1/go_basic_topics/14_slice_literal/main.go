package main

import(
	"fmt"
)

func main()  {
	
	//slices are the most common collection type
	//It is dynamic and grow
	//[]type{...}

	results :=[]string{"kishore","kumar"}
	fmt.Println(results[0])

	results[1]="dinesh"
	fmt.Println(results)

	nums :=[]int{}

	nums = append(nums,10)
	nums = append(nums, 20,30)

	fmt.Println(nums)


}