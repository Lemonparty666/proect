package main

import(
	"fmt"
	"encoding/json"
	"log"
)

type User struct{
	Id int `json:"ID`
	Name string `json:"Name"`
	Age int `json:"Age"`
}

func main(){
	var a User = User{1,"Tom",25}
	jsonData, err := json.Marshal(a)
	if err != nil {
		log.Fatalln("Ошибка создания json", err)
	}

	fmt.Println(string(jsonData))
}