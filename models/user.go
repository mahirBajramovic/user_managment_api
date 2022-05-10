package models

import (
	"database/sql"

	"github.com/mahirB/user_managment_api/services"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Status    int    `json:"status" db:"status"`
}

type UserFilter struct {
	// Uncomment to start using serverside pagination
	/* 	Limit  int    `json:"limit" db:"Limit"`
	   	Offset int    `json:"offset" db:"Offset"` */
	Status string `json:"status" db:"Status"`
}

func UsersList(filter UserFilter) (users []User, err error) {
	query := `SELECT 
				id,first_name,last_name,username,email,status
			  FROM user`

	if filter.Status != "-1" {
		query += " WHERE status = :Status"
	}

	query += " ORDER BY id ASC"
	// Uncomment to start using serverside pagination

	/* 	query += " LIMIT :Offset, :Limit"
	 */
	db := services.Access.GetDB()

	rows, err := db.NamedQuery(query, filter)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user User

		err = rows.StructScan(&user)
		if err != nil {
			return
		}

		users = append(users, user)
	}

	return users, err
}

func UserCount(filter UserFilter) (count int, err error) {
	db := services.Access.GetDB()
	query := "SELECT COUNT(*) FROM user"

	if filter.Status != "-1" {
		query += " WHERE status = ?"
		err = db.Get(&count, query, filter.Status)
	} else {
		err = db.Get(&count, query)
	}

	return
}

func (u *User) Create() error {
	db := services.Access.GetDB()

	query := `INSERT INTO user
					(first_name,last_name,username,password,email,status)
				VALUES
					(:first_name,:last_name,:username,:password,:email,:status)`

	_, err := db.NamedExec(query, u)

	return err
}

func UserDelete(id int, tx *sql.Tx) error {
	query := `DELETE FROM user WHERE id = ?`

	_, err := tx.Exec(query, id)

	return err
}

func (u *User) Get() error {
	db := services.Access.GetDB()

	query := `SELECT 
					id,first_name,last_name,username,email,status
				FROM user WHERE id = ?`

	err := db.Get(u, query, u.ID)

	return err

}

func (u *User) Update() error {
	db := services.Access.GetDB()

	query := `UPDATE user
				SET first_name=:first_name,last_name=:last_name,username=:username,
					password=:password,email=:email,status=:status
				WHERE id = :id`

	_, err := db.NamedExec(query, u)

	return err
}

func CreateTransaction() (tx *sql.Tx, err error) {
	db := services.Access.GetDB()
	tx, err = db.Begin()
	return
}
