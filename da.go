package main
import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type User struct{
	UserId int `json:"UserId"`
	Id int `json:"Id"`
	Title string `json:"TItle"`
	Completed bool `json:"Completed"`
}

func (u *User) updt(title string, completed bool){
	u.Title = title
	u.Completed = completed
}

func (u User) print(){
	fmt.Println(u.UserId, u.Id, u.Title, u.Completed)
}

func User1() (*User, error) {
	client := &http.Client{}
	req, err :=http.NewRequest( http.MethodGet,"https://jsonplaceholder.typicode.com/todos/1",nil)
	if err != nil{
		log.Fatalln(nil,err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(nil, err)
		return
	}
	defer resp.Body.Close()
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err!=nil{
		log.Fatalln(nil, err)
		return
	}
	return &user,nil
}

func main(){
	user,err :=User1()
		if err != nil{
			log.Fatalln(err)
		}
	user.print()
	user.updt(true)
	user.print()

}	

