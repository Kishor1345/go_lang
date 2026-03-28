package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	res :=map[string]any{
		"ok":true,
		"message":"JSON encode Successfully",
		"datatime":time.Now().UTC(),
	}

	_=json.NewEncoder(w).Encode(res)
}

func main()  {
	
	http.HandleFunc("/ok",successHandler)

	fmt.Println("Server Running Port : 3000")
	err := http.ListenAndServe(":3000",nil)
	fmt.Println(err)
}