package controller

import (
	"database/sql"
	"fmt"

	"github.com/Utsavch189/api_mysql/internal/config"
	"github.com/Utsavch189/api_mysql/internal/models/requests"
)

func Register(user *requests.User) (*requests.User, error) {
	db, err := config.Connect()

	if err != nil {
		return &requests.User{}, err
	}
	defer db.Close()

	query := `Insert into User(id,username,password,created_at,is_active) Values(?,?,?,?,?)`
	_, cerr := db.Exec(query, user.Id, user.UserName, user.Password, user.CreatedAt, user.IsActive)
	if cerr != nil {
		return &requests.User{}, cerr
	}

	return user, nil
}

func GetAUserByUsername(username string) (requests.User, error) {
	db, err := config.Connect()
	var user requests.User

	if err != nil {
		return user, err
	}
	defer db.Close()

	query := `Select * from User where username=?`
	row := db.QueryRow(query, username)

	err = row.Scan(
		&user.Id,
		&user.UserName,
		&user.Password,
		&user.CreatedAt,
		&user.IsActive,
	)

	if err == sql.ErrNoRows {
		return user, fmt.Errorf("user with username %s not found", username)
	}
	if err != nil {
		return user, err
	}

	return user, nil
}
