package main

import(
	"fmt"
	"encoding/json"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func (u *User) ageupdt(a int) {
	u.Age=a
}

func main(){
	a := `{"name":"Мария","age":25,"email":"maria@example.com"}`
	var b User
	err := json.Unmarshal([]byte(a),&b)
	if err!= nil{
		log.Fatalln("Ошибка сериализации", err)
	}
	fmt.Println(b)
	b.ageupdt(38)
	fmt.Println(b)

}