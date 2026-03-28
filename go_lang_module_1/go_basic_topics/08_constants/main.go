package main


import(
	"fmt"
)

func main()  {
	//fixed value not change 
	const appName = "Go Basics"

	//typed constants

	const maxUpload int =25

	const discounterPrice float64 = 52.5

	fmt.Println(appName,maxUpload,discounterPrice)
}