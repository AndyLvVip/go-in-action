package main

import "fmt"

type user struct {
	name string
	email string
}

type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("send email notification to: %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

type admin struct {
	user
	level string
}


func main() {
	ad := admin{
		user: user {
			name: "andy lv",
			email: "andy.lv.vip@hotmail.com",
		},
		level: "super",
	}
	sendNotification(&ad)
	sendNotification(&ad.user)
}