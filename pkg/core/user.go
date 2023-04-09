package core

type User struct {
	TelegramID int64
	IsAdmin    bool
}

type CreateUserParams struct {
	TelegramID int
	FirstName  string
	LastName   string
	Username   string
	LangCode   string
}
