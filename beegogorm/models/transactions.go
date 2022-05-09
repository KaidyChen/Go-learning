package models

type Bank struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Balance float64 `json:"balance"`
}

func (Bank) TableName() string {
	return "bank"
}
