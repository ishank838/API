package model

type User struct{
	User_id uint64 `json:"user_id"`
	User_name string `json:"user_name"`
	Name string `json:"name"`
	Email string `json:"email"`
}