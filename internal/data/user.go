package data

import "rss/internal/validator"

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ValidateUser(v *validator.Validator, user *User) {

	v.Check(user.Name != "", "name", "field must be provided")
	v.Check(len(user.Name) <= 200, "name", "field must not exceed 200 bytes")
}
