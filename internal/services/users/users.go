package users

import (
	db "../../database/users"
	. "../../models/users"
)

func AddUser(user *User) error {
	return db.AddUser(user)
}
