package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Account struct {
	ID            int    `json:"id"`
	AccountNumber int    `json:"accountnumber"`
	Name          string `json:"name"`
	Email         string `json:"email"`
}
