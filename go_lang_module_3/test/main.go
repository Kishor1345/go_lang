package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main()  {
	
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("Running Server : 7000")
	err :=http.ListenAndServe(":7000",nil)
	fmt.Println(err)
}

func writeJSON(w http.ResponseWriter,status int,data any){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	_=json.NewEncoder(w).Encode(data)
}



func helloHandler(w http.ResponseWriter,r *http.Request)  {
	if r.Method != http.MethodGet{
		writeJSON(w,http.StatusBadRequest,map[string]any{
			"ok":false,
			"message":"Invalid json fomrat",
		})
		return
	}

	writeJSON(w,http.StatusAccepted,map[string]any{
		"ok":true,
		"message":"success",
	})
}