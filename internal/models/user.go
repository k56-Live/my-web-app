package models

type User struct {
	ID        int
	Username  string
	Email     string
	FirstName string
	LastName  string
}

func NewUser(id int, username, email, firstName, lastName string) *User {
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
