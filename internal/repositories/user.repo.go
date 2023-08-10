package repositories

import (
	"fmt"
	"log"
	"ninja1cak/coffeshop-be/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `INSERT INTO public.user (
		email, 
		password, 
		phone_number
		)
	VALUES(
		:email, 
		:password, 
		:phone_number		
	)`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "user success created", nil
}

func (r *RepoUser) GetUser() (interface{}, error) {
	query := `SELECT 
		email,
		first_name,
		last_name,
		phone_number,
		address,
		birth_date
	FROM
		public.user`

	data := []models.User{}

	err := r.Select(&data, query)
	log.Println(data)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (r *RepoUser) UpdateUser(data *models.User) (string, error) {
	set := ""

	if data.First_name != nil {
		set += "first_name = :first_name,"
	}

	if data.Last_name != nil {
		set += "last_name = :last_name,"
	}

	if data.Address != nil {
		set += "address = :address,"
	}

	if data.Birth_date != nil {
		set += "birth_date = :birth_date,"
	}

	if data.Photo_profile != nil {
		set += "photo_profile = :photo_profile,"
	}

	set += "updated_at = NOW()"

	query := fmt.Sprintf(`UPDATE public.user
	SET
		%s
	WHERE
		email = :email
		`, set)

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "update user success", nil
}

func (r *RepoUser) DeleteUser(data *models.User) (string, error) {
	query := `DELETE FROM public.user
	WHERE
		email = :email
		`

	res, err := r.NamedExec(query, data)
	rowsChange, err := res.RowsAffected()
	if rowsChange == 0 {
		return "data not found", nil
	}
	if err != nil {
		return "", err
	}

	return "Delete user success", nil
}
