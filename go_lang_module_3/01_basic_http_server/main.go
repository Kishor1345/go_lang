package main

import (
	"fmt"
	"net/http"
)

//creating function for the request and response 
func helloHandler(w http.ResponseWriter,r *http.Request){

	if r.Method != http.MethodGet{
		http.Error(w,"Only Get is Allowed",http.StatusMethodNotAllowed)
		return
	}

	//write function only accept byte slice
	_, _ = w.Write([]byte("Hello from Go net/http server"))
}

func main()  {
	
	//creating register router
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("try going to 3000 port")

	//creating server to listen and port declaration
	err :=http.ListenAndServe(":3000",nil)
	fmt.Println(err)
}