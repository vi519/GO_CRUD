package models

type Account struct {
	ID            int    `json:"id"`
	AccountNumber int    `json:"accountnumber"`
	Name          string `json:"name"`
	Email         string `json:"email"`
}

func (Account) TableName() string {
	return "accounts"
}
