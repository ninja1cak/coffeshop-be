package models

import (
	"time"
)

type User struct {
	User_id       string     `db:"user_id" form:"user_id" valid:"-"`
	Email         string     `db:"email" form:"email" valid:"email"`
	Password      string     `db:"password" form:"password" valid:"minstringlength(4)~Password minimal 4"`
	First_name    *string    `db:"first_name" form:"first_name" valid:"-"`
	Last_name     *string    `db:"last_name" form:"last_name" valid:"-"`
	Phone_number  string     `db:"phone_number" form:"phone_number" valid:"-"`
	Address       *string    `db:"address" form:"address" valid:"-"`
	Birth_date    *time.Time `db:"birth_date" form:"birth_date" valid:"-"`
	Gender        *string    `db:"gender" form:"gender" valid:"-"`
	Photo_profile *string    `db:"photo_profile" form:"photo_profile" valid:"-"`
	Status        *string    `db:"status" form:"status" valid:"-"`
	Created_at    string     `db:"created_at" form:"created_at" valid:"-"`
	Updated_at    *string    `db:"updated_at" form:"updated_at" valid:"-"`
	Role          string     `db:"role" valid:"-"`
	// File          *multipart.FileHeader `form:"file" binding:"required"`
}
