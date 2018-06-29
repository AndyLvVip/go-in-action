package main

import "fmt"

type user struct {
	name string
	email string
}

type notifier interface {
	notify()
}

type admin struct {
	user
	level string
}

func (u *user) notify() {
	fmt.Printf("send user notification to: %s<%s>\n", u.email, u.name)
}

func (ad *admin) notify() {
	fmt.Printf("send admin notification to: %s<%s>\n", ad.name, ad.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user : user {
			name : "andy lv",
			email: "andy.lv.vip@hotmail.com",
		},
		level: "super",
	}

	ad.notify()
	ad.user.notify()
	sendNotification(&ad)
	sendNotification(&ad.user)
}
