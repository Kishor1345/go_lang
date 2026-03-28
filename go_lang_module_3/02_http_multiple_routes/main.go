package main

import (
	"net/http"
	"fmt"
)

func rootHandler(w http.ResponseWriter,r *http.Request){
	_,_ = w.Write([]byte("Welcome try to /hello?name=kishore"))
}

func hellohandler(w http.ResponseWriter,r * http.Request){
	name :=r.URL.Query().Get("name")

	if name == ""{
		name = "Guest"
	}

	_,_=w.Write([]byte("Hello Welcome "+name))
}

func main() {

	http.HandleFunc("/",rootHandler)
	http.HandleFunc("/hello",hellohandler)
	fmt.Println("try going to 3000 port")
	
	err :=http.ListenAndServe(":3000",nil)
	
	fmt.Println(err)
}