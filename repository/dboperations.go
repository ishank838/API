package repository

import (
	"API/model"
	"database/sql"
	"log"
)

func CreateUser(db *sql.DB, u *model.User) (uint64, string) {

	var userID uint64
	var rowCount int32

	if len(u.User_name) == 0 {
		return 0, "Username can't Be nil"
	}

	if len(u.Email) == 0 {
		return 0, "Email can't be nil"
	}

	err := db.QueryRow("select count(user_id) from users where username=$1 or email=$2",
		u.User_name,
		u.Email).Scan(&rowCount)

	if err != nil {
		log.Fatal(err)
	}

	if rowCount > 0 {
		return 0, "User already exists"
	} else {

		err = db.QueryRow("insert into users values(DEFAULT,$1,$2,$3) RETURNING user_id;",
			u.User_name,
			u.Name,
			u.Email).Scan(&userID)

		if err != nil {
			return 0, "Inconsistent Data"
			panic(err)
		}

		u.User_id = userID
		return userID, "Sucessfull"
	}
}

func UpdateUser(db *sql.DB, id uint64, u model.User) (string, *model.User) {
	var rowCount uint64

	err := db.QueryRow("select count(user_id) from users where user_id=$1;", id).Scan(&rowCount)

	if err != nil {
		panic(err)
	}
	if rowCount == 0 {
		return "ID doesn't exist", nil
	} else {

		_, err := db.Exec("update users set username=$1, name=$2, email=$3 where user_id=$4;",
			u.User_name,
			u.Name,
			u.Email,
			id)

		if err != nil {
			panic(err)
		}

		u.User_id = id
		return "Sucessfull", &u
	}
}

func DeleteUser(db *sql.DB, id uint64) string {

	var rowCount uint64

	err := db.QueryRow("Select count(user_id) from users where user_id=$1;", id).Scan(&rowCount)

	if err != nil {
		panic(err)
	}

	if rowCount == 0 {
		return "ID doesn't exist"
	} else {
		_, err := db.Exec("delete from users where user_id=$1",id)

		if err != nil {
			panic(err)
		}
		return "Sucessfull"
	}
}

func ListUsers(db *sql.DB, limit uint64, offset uint64) []model.User {

	r, err := db.Query("Select * from users ORDER BY user_id LIMIT $1 OFFSET $2",
		limit,
		offset)

	if err != nil {
		panic(err)
	}

	var user model.User
	var users []model.User
	for r.Next() {
		err := r.Scan(&user.User_id, &user.User_name, &user.Name, &user.Email)

		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}
