package partie1

import "fmt"

type User struct {
	Id 			int 	`json:"id"`
	FirstName 	string 	`json:"first_name"`
	Subscribed 	bool 	`json:"subscribed"`
}

type Polite interface {
	Greeting() string
}

func (u User) Greeting() string {
	return "Coucou"

}

func SayHello(p Polite) {
	fmt.Println(p.Greeting())
}


func Partie1() {
	u := User{}
	SayHello(u)
}