package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)
func writeJSON(w http.ResponseWriter,status int,data any){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	_=json.NewEncoder(w).Encode(data)
}

type TestRequest struct{
	Name string `json:"name"`
}

func testHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		writeJSON(w,http.StatusMethodNotAllowed,map[string]any{
			"ok":false,
			"message":"Only Post is Allowed",
		})
		return
	}
	defer r.Body.Close()

	var req TestRequest
	err:=json.NewDecoder(r.Body).Decode(&req)
	if err !=nil{
		writeJSON(w,http.StatusBadRequest,map[string]any{
			"ok":false,
			"message":"Invalid json fomrat",
		})
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == ""{
		writeJSON(w,http.StatusBadRequest,map[string]any{
			"ok":false,
			"error":"name must be not be empty",
		})
		return
	}

	writeJSON(w,http.StatusAccepted,map[string]any{
		"ok":"success",
		"data":req,
		"time":time.Now().UTC(),
	})
}

func main()  {
	
	http.HandleFunc("/test",testHandler)
	fmt.Println("Server Running Port : 3000")

	err :=http.ListenAndServe(":3000",nil)
	fmt.Println(err)


}