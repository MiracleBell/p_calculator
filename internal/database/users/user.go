package users

import (
	. ".."
	. "../../models/users"
	"log"
)

var db, _ = OpenConnection()

func AddUser(user *User) error {
	query := "INSERT INTO users(login, password, email, created_at, last_updated_at) VALUES(?,?,?,?,?)"

	result, err := db.Exec(query,
		user.Login,
		user.Password,
		user.Email,
		user.CreatedAt,
		user.LastUpdatedAt)

	if err != nil {
		log.Fatal("Can't add user in DB")
		return err
	}
	id, _ := result.LastInsertId()
	user.ID = uint64(id)
	return nil
}
