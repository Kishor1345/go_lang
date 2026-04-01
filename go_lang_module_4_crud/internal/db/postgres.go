package db

import (
	"crud/internal/config"
	"database/sql"
	"fmt"
)

func Connect(cfg config.Config)(*sql.DB,error){
	connectStr:=fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode-disable",
		cfg.Host,cfg.Port,cfg.Username,cfg.Password,cfg.Dbname,
	)

	db,err:=sql.Open("postgres",connectStr)
	if err!=nil{
		return nil,fmt.Errorf("db connection failed")
	}

	if err:=db.Ping();err!=nil{
		return nil,err
	}
	fmt.Println("✅ Database Connected")
	
	return db,nil
}