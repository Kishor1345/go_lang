package main

import (
	"crud/internal/config"
	"crud/internal/db"
	"crud/internal/server"
	"fmt"
	"log"
)

//config ->db ->router -> run server

func main()  {
	cfg,err:=config.Load()
	fmt.Println(err)
	if err!=nil{
		log.Fatalf("Config error")
	}

	db,err :=db.Connect(cfg)
	if err !=nil{
		log.Fatalf("DB error")
	}
	fmt.Println(db)

	router :=server.NewRouter()
	server_port:="3000"
	addr :=fmt.Sprintf(":%s",server_port)
	if err:=router.Run(addr);err !=nil{
		log.Fatalf("server failed")
	}
}