package models

import "fmt"

type User struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Wallet *Wallet `json:"wallet"`
}

func (u *User) String() string {
	return fmt.Sprintf("User id: %d, name: %s", u.Id, u.Name)
}
