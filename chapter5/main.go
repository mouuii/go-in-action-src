package  main
import (
	"fmt"
)
type notifier interface {
	notify()
}
type user struct {
	name string
	email string
}

func (u *user)notify(){
	fmt.Printf("sending user email to %s<%s>",u.name,u.email)
}

func (u *admin)notify(){
	fmt.Printf("sending user email to %s<%s>",u.name,u.email)
}

type admin struct {
	user
	level string
	name string
	email string
}

func main(){
	ad := admin{
		user:user{
			name:"jj",
			email:"xx",
		},
		level:"super",
		name:"dddddddddd",
		email:"xxxxxxxxxxxxxxxxxxx",
	}
	sendNotify(&ad)
}

func sendNotify(n notifier){
	n.notify()
}