package core

type User struct {
	ID int `json:"id"` // Telegram user ID
}

type CreateUserParams struct {
	ID   int    `json:"id"`   // Telegram user ID
	Lang string `json:"lang"` // Language code
}
