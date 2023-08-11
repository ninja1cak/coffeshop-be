package repositories

import (
	"errors"
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
		phone_number,
		role
		)
	VALUES(
		:email, 
		:password, 
		:phone_number,
		:role
	)`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "user success created", nil
}

func (r *RepoUser) GetUser(user_id string) (interface{}, error) {
	var queryId string = ""
	var err error
	if user_id != "" {
		queryId = "WHERE user_id = $1"
	}

	query := fmt.Sprintf(`SELECT 
	email,
	first_name,
	last_name,
	phone_number,
	address,
	birth_date
FROM
	public.user %s`, queryId)

	data := []models.User{}

	if user_id == "" {

		if err = r.Select(&data, query); err != nil {
			return "", err
		}
	} else {
		if err = r.Select(&data, query, user_id); err != nil {
			return "", err
		}
	}
	log.Println(data)

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

	if *data.Photo_profile != "" {
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

func (r *RepoUser) GetAuthData(email string) (*models.User, error) {
	var result models.User
	query := `SELECT user_id, email, password, role from public.user where email = ?`

	if err := r.Get(&result, r.Rebind(query), email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil

}
