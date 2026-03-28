package main

import(
	"fmt"
)



func main()  {
	
	var city string
	city = "London"

	//inferred to string
	var channel = "kishore"

	// :=
	subject := "learning"
	subject = subject+"english"

	//multiple variable in same line
	likes ,comment :=100,30

	fmt.Println(city,channel,subject,likes,comment)
}