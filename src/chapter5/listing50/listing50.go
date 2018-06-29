package main

import "fmt"

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("send email to notify %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name: "andy lv",
			email: "andy.lv.vip@hotmail.com",
		},
		level: "super",
	}

	ad.user.notify()
	ad.notify()
}
