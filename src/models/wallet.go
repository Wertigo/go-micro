package models

type Wallet struct {
	Id      int     `json:"id"`
	User_id int     `json:"-"`
	Money   float64 `json:"money"`
}
