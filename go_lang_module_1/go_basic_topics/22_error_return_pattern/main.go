package main

import (
	"fmt"
	"log"
	"strconv"
)

func main()  {
	
	//go dont use exception for normal failures
	// functions ->return error as normal return values

	//val,err :=something()
	//if err ! = nill {handle the error}

	if err:=run();err !=nil{
		//give timestamps with error while using log.fatal
		log.Fatal(err)
	}


}

func run() error{
	input :="30"
	level,err :=parseLevel(input)
	if err!=nil{
		return err
	}
	fmt.Println(level)
	return err
}

func parseLevel(s string) (int,error){
	//(values,err)return pattern
	//nil error means no error
	//non nil means error

	//pattern
	//Atio convert string to int
	n,err:=strconv.Atoi(s)
	if err !=nil{
		return 0,fmt.Errorf("level must be a number")
	}

	if n<1 || n>5{
		return 0,fmt.Errorf("level must be 1 and 5")
	}

	return n,nil
}