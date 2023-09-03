package models

import (
	"time"
)

type User struct {
	User_id       string     `db:"user_id" form:"user_id" json:"user_id,omitempty" valid:"-"`
	Email         string     `db:"email" form:"email" json:"email,omitempty" valid:"email"`
	Password      string     `db:"password" form:"password" json:"password,omitempty"  valid:"minstringlength(4)~Password minimal 4"`
	First_name    *string    `db:"first_name" form:"first_name" json:"first_name,omitempty"  valid:"-"`
	Last_name     *string    `db:"last_name" form:"last_name" json:"last_name,omitempty" valid:"-"`
	Phone_number  string     `db:"phone_number" form:"phone_number" json:"phone_number,omitempty" valid:"-"`
	Address       *string    `db:"address" form:"address" json:"address,omitempty" valid:"-"`
	Birth_date    *time.Time `db:"birth_date" form:"birth_date" json:"birth_date,omitempty" valid:"-"`
	Gender        *string    `db:"gender" form:"gender" json:"gender,omitempty" valid:"-"`
	Photo_profile *string    `db:"photo_profile" form:"photo_profile" json:"photo_profile,omitempty" valid:"-"`
	Status        *string    `db:"status" form:"status" json:"status,omitempty" valid:"-"`
	Created_at    string     `db:"created_at" form:"created_at" json:"created_at,omitempty" valid:"-"`
	Updated_at    *string    `db:"updated_at" form:"updated_at" json:"updated_at,omitempty"  valid:"-"`
	Role          string     `db:"role" valid:"-" json:"role,omitempty"`
	// File          *multipart.FileHeader `form:"file" binding:"required"`
}
