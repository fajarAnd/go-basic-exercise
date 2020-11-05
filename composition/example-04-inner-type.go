package main

import "fmt"

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Println("Ada Email untuk :", u.name, " : ", u.email)
}

type admin struct {
	user
	level string
}

func main()  {
	ad := admin{
		user: user{
			name: "fajar",
			email: "fajar@gmail",
		},
		level: "super",
	}
	ad.notify()
	ad.user.notify()
}