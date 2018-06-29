package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("send notification tp %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user {"andy", "andy.lv.vip@hotmail.com"}

	u.notify()

	sendNotification(&u)
}