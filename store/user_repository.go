package store

import (
	"log"

	"github.com/maksimartemev/golang-db-pg-example/model"
	"golang.org/x/crypto/bcrypt"
)

func GetById(userId int) (*model.User, error) {
	db := GetDB()
	defer db.Close()

	user := &model.User{}
	err := db.QueryRow("SELECT id, email FROM users WHERE id = $1", userId).Scan(&user.ID, &user.Email)

	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func Create(email string, password string) (*model.User, error) {
	db := GetDB()
	defer db.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return &model.User{}, err
	}

	user := &model.User{}
	errr := db.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email", email, string(hash)).Scan(&user.ID, &user.Email)

	if errr != nil {
		return &model.User{}, errr
	}

	return user, nil
}

func List() ([]model.User, error) {
	db := GetDB()
	defer db.Close()

	var users []model.User

	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)

		if err = rows.Err(); err != nil {
			return users, err
		}
	}
	return users, nil
}
