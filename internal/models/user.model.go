package models

import (
	"time"
)

type User struct {
	User_id       string     `db:"user_id" form:"user_id"`
	Email         string     `db:"email" form:"email"`
	Password      string     `db:"password" form:"password"`
	First_name    *string    `db:"first_name" form:"first_name"`
	Last_name     *string    `db:"last_name" form:"last_name"`
	Phone_number  string     `db:"phone_number" form:"phone_number"`
	Address       *string    `db:"address" form:"address"`
	Birth_date    *time.Time `db:"birth_date" form:"birth_date"`
	Gender        *string    `db:"gender" form:"gender"`
	Photo_profile *string    `db:"photo_profile" form:"photo_profile"`
	Status        *string    `db:"status" form:"status"`
	Created_at    string     `db:"created_at" form:"created_at"`
	Updated_at    *string    `db:"updated_at" form:"updated_at"`
	// File          *multipart.FileHeader `form:"file" binding:"required"`
}
