package config

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Username string 
	Password string
	Host string
	Port string
	Dbname string
}

func Load()(Config,error){

	//godotenv.Load() reads .env and sets them into the process env
	//os.getenv ->reads those values

	if err :=godotenv.Load(".env");err !=nil{
		fmt.Println(err)
		return Config{},fmt.Errorf("Failed to load .env")
	}
	
	username,err:=extractEnv("USER_NAME")
	if err!=nil{
		return Config{},err
	}

	password,err:=extractEnv("PASSWORD")
	if err!=nil{
		return Config{},err
	}

	host,err:=extractEnv("HOST")
	if err!=nil{
		return Config{},err
	}

	port,err:=extractEnv("PORT")
	if err!=nil{
		return Config{},err
	}

	dbname,err:=extractEnv("DBNAME")
	if err!=nil{
		return Config{},err
	}

	config:=Config{
		Username:username,
		Password: password,
		Host: host,
		Port: port,
		Dbname: dbname,
	}
	return config,nil
}

func extractEnv(key string)(string,error){

	val :=os.Getenv(key)

	if val == ""{
		return "",fmt.Errorf("missing req key")
	}

	return val,nil
}