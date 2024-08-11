package models

import (
	"ankumar/events-api/db"
	"ankumar/events-api/utils"
	"errors"
)

type User struct {
	ID         int64  `binding: "required"`
	Email      string `binding: "required"`
	Password   string `binding: "required"`
	IdentityId string
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password, identityId) values(?, ?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.IdentityId = utils.GenerateIdentityId()

	result, err := stmt.Exec(u.Email, hashedPassword, u.IdentityId)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT password, identityId FROM users WHERE email=?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	row.Scan(&retrievedPassword, &u.IdentityId)

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}

func GetUserByIdentityId(identityId string) (*User, error) {
	query := "SELECT id FROM users WHERE identityId = ?"

	row := db.DB.QueryRow(query, identityId)

	var user User
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
