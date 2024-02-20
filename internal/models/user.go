package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsAdmin   bool
}

func NewUser(id int, isAdmin bool, firstName, lastName, email, password string) *User {
	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		IsAdmin:   isAdmin,
	}
}
