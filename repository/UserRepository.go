package repository

import (
	"blog/model"
	"database/sql"
	"errors"
	// "github.com/pelletier/go-toml/query"
)

func GetAllUsers(db *sql.DB, page int, limit int) ([]model.User, int, error) {
	// check the limitedd
	total, err := countTotalUsers(db)
	if err != nil {
		return nil, 0, err
	}

	if page > total {
		return nil, 0, errors.New("page out of range")
	}

	// get db statement

	query := `select u."userId", u."userName", u."password", u."role", u."createdAt", u."updatedAt" from "user" u where u."deletedAt" is null limit 9 offset 3`
	result, err := db.Query(query)
	if err != nil {
		return nil, 0, err
	}

	// binding the data from db to object

	slicesUsers := make([]model.User, 0)
	for result.Next() {
		var user model.User
		err = result.Scan(&user.UserId, &user.UserName, &user.Password, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			return nil, 0, err
		}
		slicesUsers = append(slicesUsers, user)
	}

	// return

	return slicesUsers, total, nil
}

func countTotalUsers(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("select count(*) from \"user\" u where u.\"deletedAt\" is null").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func AddUser(db *sql.DB, user model.User) (int64, error) {

	if user.Role == "" {
		query := `INSERT INTO "user" ("userName", "password",) VALUES ($1, $2)`
		result, err := db.Exec(query, user.UserName, user.Password)
		if err != nil {
			return 0, err
		}
		return result.RowsAffected()
	} else {
		query := `INSERT INTO "user" ("userName", "password", "role") VALUES ($1, $2, $3)`

		result, err := db.Exec(query, user.UserName, user.Password, user.Role)
		if err != nil {

			return 0, err
		}
		return result.RowsAffected()
	}

}

func GetAdmin(db *sql.DB) ([]model.User, error) {
	query := `select u."userId", u."userName", u."password",  u."createdAt", u."updatedAt" from "user" u where u."deletedAt" is null and u."role" = 'admin'`
	result, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	slicesUsers := make([]model.User, 0)
	for result.Next() {
		var user model.User
		err = result.Scan(&user.UserId, &user.UserName, &user.Password, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			return nil, err
		}
		slicesUsers = append(slicesUsers, user)
	}
	return slicesUsers, nil
}
